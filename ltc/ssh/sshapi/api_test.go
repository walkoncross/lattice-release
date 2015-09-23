package sshapi_test

import (
	"bytes"
	"errors"
	"io"
	"net"
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry-incubator/lattice/ltc/sshapi"
	"github.com/cloudfoundry-incubator/lattice/ltc/sshapi/fake_ssh_client"
	"golang.org/x/crypto/ssh"
)

var _ = Describe("SecureDialer", func() {
	var secureDialer *ssh.SecureDialer

	BeforeEach(func() {
		secureDialer = &ssh.SecureDialer{}
	})

	Describe("#Dial", func() {
		It("passes correct args to the Dial impl", func() {
			dialCalled := false
			secureDialer.DialFunc = func(network, addr string, config *ssh.ClientConfig) (*ssh.Client, error) {
				Expect(network).To(Equal("tcp"))
				Expect(addr).To(Equal("10.0.12.34:2222"))
				Expect(config.User).To(Equal("diego:app-name/2"))

				Expect(config.Auth).To(HaveLen(1))

				actualSecret := reflect.ValueOf(config.Auth[0]).Call([]reflect.Value{})[0].Interface()
				Expect(actualSecret).To(Equal("user:past"))

				dialCalled = true

				return nil, errors.New("")
			}

			secureDialer.Dial("diego:app-name/2", "user", "past", "10.0.12.34:2222")

			Expect(dialCalled).To(BeTrue())
		})
	})
})

type mockConn struct {
	io.Reader
	io.Writer
	nilNetConn
	closed bool
}

type nilNetConn struct {
	net.Conn
}

func (m *mockConn) Close() error {
	m.closed = true
	return nil
}

var _ = Describe("SecureClient", func() {
	var (
		secureClient  *ssh.SecureClient
		mockSSHClient *fake_ssh_client.FakeSSHClient
	)

	BeforeEach(func() {
		mockSSHClient = &fake_ssh_client.FakeSSHClient{}
		secureClient = &ssh.SecureClient{mockSSHClient}
	})

	Describe("#Accept", func() {
		It("should dial a remote connection", func() {
			localConn := &mockConn{Reader: &bytes.Buffer{}, Writer: &bytes.Buffer{}}
			remoteConn := &mockConn{Reader: &bytes.Buffer{}, Writer: &bytes.Buffer{}}
			mockSSHClient.DialReturns(remoteConn, nil)

			Expect(secureClient.Accept(localConn, "some remote address")).To(Succeed())

			Expect(mockSSHClient.DialCallCount()).To(Equal(1))
			protocol, address := mockSSHClient.DialArgsForCall(0)
			Expect(protocol).To(Equal("tcp"))
			Expect(address).To(Equal("some remote address"))
		})

		It("should copy data in both directions", func() {
			localConnBuffer := &bytes.Buffer{}
			remoteConnBuffer := &bytes.Buffer{}
			localConn := &mockConn{Reader: bytes.NewBufferString("some local data"), Writer: localConnBuffer}
			remoteConn := &mockConn{Reader: bytes.NewBufferString("some remote data"), Writer: remoteConnBuffer}
			mockSSHClient.DialReturns(remoteConn, nil)

			Expect(secureClient.Accept(localConn, "some remote address")).To(Succeed())

			Expect(localConn.closed).To(BeTrue())
			Expect(remoteConn.closed).To(BeTrue())
			Expect(localConnBuffer.String()).To(Equal("some remote data"))
			Expect(remoteConnBuffer.String()).To(Equal("some local data"))
		})

		Context("when dialing a remote connection fails", func() {
			It("should return an error", func() {
				mockSSHClient.DialReturns(nil, errors.New("some error"))
				err := secureClient.Accept(nil, "some remote address")
				Expect(err).To(MatchError("some error"))
			})
		})
	})

	Describe("#Open", func() {
		It("should ", func() {
			mockSessionStdin = &bytes.Buffer{}
			mockSessionStdout = bytes.NewBufferString("some session out data")
			mockSessionStderr = bytes.NewBufferString("some session err data")
			mockSession := &fake_session.FakeSession{}
			mockSessionFactory.NewReturns(mockSession, nil)
			mockSession.StdinPipeReturns(mockSessionStdin, nil)
			mockSession.StdoutPipeReturns(mockSessionStdout, nil)
			mockSession.StderrPipeReturns(mockSessionStderr, nil)

			Expect(secureClient.Open(100, 200)).To(Succeed())

			Expect(mockSession.RequestPtyCallCount()).To(Equal(1))
			termType, height, width, modes := fakeSession.RequestPtyArgsForCall(0)
			Expect(termType).To(Equal("defaultterm"))
			Expect(height).To(Equal(200))
			Expect(width).To(Equal(100))
			Expect(modes[ssh.ECHO]).To(Equal(1))
			Expect(modes[ssh.TTY_OP_ISPEED]).To(Equal(115200))
			Expect(modes[ssh.TTY_OP_OSPEED]).To(Equal(115200))
		})
	})

	Describe("#Resize", func() {

	})

	Describe("#KeepAlive", func() {

	})
})