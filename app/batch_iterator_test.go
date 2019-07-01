package app

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var iteratorTestData = `{"CODE_FILE":"../src/core/job.c","SYSLOG_FACILITY":"3","_SYSTEMD_UNIT":"init.scope","_SYSTEMD_SLICE":"-.slice","JOB_ID":"158","_CMDLINE":"/sbin/init splash","_COMM":"systemd","JOB_TYPE":"start","_BOOT_ID":"6ad27b41a7994a9592b7a9f031b4748e","_GID":"0","_SYSTEMD_CGROUP":"/init.scope","MESSAGE":"Starting Docker Application Container Engine...","__REALTIME_TIMESTAMP":"1562000012079123","_TRANSPORT":"journal","UNIT":"docker.service","_SOURCE_REALTIME_TIMESTAMP":"1562000012077797","CODE_FUNC":"job_log_begin_status_message","_EXE":"/usr/lib/systemd/systemd","_SELINUX_CONTEXT":"unconfined\n","__CURSOR":"s=f9eccc9100f5452ab6e7af6c221c296b;i=39f;b=6ad27b41a7994a9592b7a9f031b4748e;m=26cce4b;t=58ca177dbf013;x=d1186ad61220bfd7","_CAP_EFFECTIVE":"3fffffffff","_MACHINE_ID":"03a7024e57114636946d8fc3bf8f4863","PRIORITY":"6","_HOSTNAME":"jarek-VirtualBox","SYSLOG_IDENTIFIER":"systemd","__MONOTONIC_TIMESTAMP":"40685131","_UID":"0","CODE_LINE":"600","INVOCATION_ID":"b16ab4b640fd4721a1d56fbfc7db0941","MESSAGE_ID":"7d4958e842da4a758f6c1cdc7b36dcc5","_PID":"1"}
{"_PID":"861","_CMDLINE":"/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock","_SYSTEMD_CGROUP":"/system.slice/docker.service","__CURSOR":"s=f9eccc9100f5452ab6e7af6c221c296b;i=3e5;b=6ad27b41a7994a9592b7a9f031b4748e;m=3529308;t=58ca178c1b4d0;x=44fcf04df65c9290","_UID":"0","_COMM":"dockerd","MESSAGE":"time=\"2019-07-01T18:53:47.137134028+02:00\" level=info msg=\"systemd-resolved is running, so using resolvconf: /run/systemd/resolve/resolv.conf\"","SYSLOG_FACILITY":"3","_SELINUX_CONTEXT":"unconfined\n","PRIORITY":"6","_SYSTEMD_SLICE":"system.slice","_CAP_EFFECTIVE":"3fffffffff","_MACHINE_ID":"03a7024e57114636946d8fc3bf8f4863","_BOOT_ID":"6ad27b41a7994a9592b7a9f031b4748e","_SYSTEMD_UNIT":"docker.service","_EXE":"/usr/bin/dockerd-ce","_TRANSPORT":"stdout","_HOSTNAME":"jarek-VirtualBox","_GID":"0","__MONOTONIC_TIMESTAMP":"55743240","_STREAM_ID":"91dd1df6941e465d8f205115fe436e36","__REALTIME_TIMESTAMP":"1562000027137232","SYSLOG_IDENTIFIER":"dockerd","_SYSTEMD_INVOCATION_ID":"b16ab4b640fd4721a1d56fbfc7db0941"}
{"_GID":"0","SYSLOG_FACILITY":"3","_COMM":"dockerd","_CMDLINE":"/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock","_PID":"861","_EXE":"/usr/bin/dockerd-ce","_SYSTEMD_INVOCATION_ID":"b16ab4b640fd4721a1d56fbfc7db0941","_UID":"0","MESSAGE":"time=\"2019-07-01T18:53:47.370438693+02:00\" level=info msg=\"parsed scheme: \\\"unix\\\"\" module=grpc","__MONOTONIC_TIMESTAMP":"55976495","SYSLOG_IDENTIFIER":"dockerd","_CAP_EFFECTIVE":"3fffffffff","_SELINUX_CONTEXT":"unconfined\n","_SYSTEMD_CGROUP":"/system.slice/docker.service","_HOSTNAME":"jarek-VirtualBox","__CURSOR":"s=f9eccc9100f5452ab6e7af6c221c296b;i=3ea;b=6ad27b41a7994a9592b7a9f031b4748e;m=356222f;t=58ca178c543f7;x=c656502506fdc4b3","_SYSTEMD_SLICE":"system.slice","_TRANSPORT":"stdout","__REALTIME_TIMESTAMP":"1562000027370487","_BOOT_ID":"6ad27b41a7994a9592b7a9f031b4748e","_MACHINE_ID":"03a7024e57114636946d8fc3bf8f4863","_SYSTEMD_UNIT":"docker.service","_STREAM_ID":"91dd1df6941e465d8f205115fe436e36","PRIORITY":"6"}
{"SYSLOG_FACILITY":"3","_GID":"0","_PID":"861","_SYSTEMD_CGROUP":"/system.slice/docker.service","MESSAGE":"time=\"2019-07-01T18:53:47.370874315+02:00\" level=info msg=\"scheme \\\"unix\\\" not registered, fallback to default scheme\" module=grpc","PRIORITY":"6","_BOOT_ID":"6ad27b41a7994a9592b7a9f031b4748e","__CURSOR":"s=f9eccc9100f5452ab6e7af6c221c296b;i=3eb;b=6ad27b41a7994a9592b7a9f031b4748e;m=35623d0;t=58ca178c54598;x=a7303fcaa2c068af","__MONOTONIC_TIMESTAMP":"55976912","_MACHINE_ID":"03a7024e57114636946d8fc3bf8f4863","_HOSTNAME":"jarek-VirtualBox","_SYSTEMD_UNIT":"docker.service","_UID":"0","_COMM":"dockerd","_EXE":"/usr/bin/dockerd-ce","_SELINUX_CONTEXT":"unconfined\n","__REALTIME_TIMESTAMP":"1562000027370904","_SYSTEMD_SLICE":"system.slice","_CAP_EFFECTIVE":"3fffffffff","SYSLOG_IDENTIFIER":"dockerd","_STREAM_ID":"91dd1df6941e465d8f205115fe436e36","_SYSTEMD_INVOCATION_ID":"b16ab4b640fd4721a1d56fbfc7db0941","_CMDLINE":"/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock","_TRANSPORT":"stdout"}
{"_STREAM_ID":"91dd1df6941e465d8f205115fe436e36","_MACHINE_ID":"03a7024e57114636946d8fc3bf8f4863","__MONOTONIC_TIMESTAMP":"55977345","PRIORITY":"6","SYSLOG_IDENTIFIER":"dockerd","_SYSTEMD_CGROUP":"/system.slice/docker.service","__REALTIME_TIMESTAMP":"1562000027371338","_EXE":"/usr/bin/dockerd-ce","_BOOT_ID":"6ad27b41a7994a9592b7a9f031b4748e","_SYSTEMD_INVOCATION_ID":"b16ab4b640fd4721a1d56fbfc7db0941","_GID":"0","_CAP_EFFECTIVE":"3fffffffff","_COMM":"dockerd","_TRANSPORT":"stdout","_PID":"861","MESSAGE":"time=\"2019-07-01T18:53:47.371298989+02:00\" level=info msg=\"ccResolverWrapper: sending new addresses to cc: [{unix:///run/containerd/containerd.sock 0  <nil>}]\" module=grpc","SYSLOG_FACILITY":"3","_SYSTEMD_SLICE":"system.slice","_UID":"0","_SELINUX_CONTEXT":"unconfined\n","_HOSTNAME":"jarek-VirtualBox","_CMDLINE":"/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock","__CURSOR":"s=f9eccc9100f5452ab6e7af6c221c296b;i=3ec;b=6ad27b41a7994a9592b7a9f031b4748e;m=3562581;t=58ca178c5474a;x=aa7477b44831e85a","_SYSTEMD_UNIT":"docker.service"}
{"_STREAM_ID":"91dd1df6941e465d8f205115fe436e36","_SYSTEMD_UNIT":"docker.service","_GID":"0","SYSLOG_IDENTIFIER":"dockerd","_CMDLINE":"/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock","_SYSTEMD_CGROUP":"/system.slice/docker.service","PRIORITY":"6","_PID":"861","__CURSOR":"s=f9eccc9100f5452ab6e7af6c221c296b;i=3ed;b=6ad27b41a7994a9592b7a9f031b4748e;m=3562634;t=58ca178c547fc;x=508b221df2b81530","_CAP_EFFECTIVE":"3fffffffff","_SELINUX_CONTEXT":"unconfined\n","_SYSTEMD_INVOCATION_ID":"b16ab4b640fd4721a1d56fbfc7db0941","_HOSTNAME":"jarek-VirtualBox","_UID":"0","__REALTIME_TIMESTAMP":"1562000027371516","__MONOTONIC_TIMESTAMP":"55977524","_EXE":"/usr/bin/dockerd-ce","_SYSTEMD_SLICE":"system.slice","SYSLOG_FACILITY":"3","_BOOT_ID":"6ad27b41a7994a9592b7a9f031b4748e","MESSAGE":"time=\"2019-07-01T18:53:47.371486742+02:00\" level=info msg=\"ClientConn switching balancer to \\\"pick_first\\\"\" module=grpc","_COMM":"dockerd","_MACHINE_ID":"03a7024e57114636946d8fc3bf8f4863","_TRANSPORT":"stdout"}
{"_UID":"0","_COMM":"dockerd","_CAP_EFFECTIVE":"3fffffffff","_SYSTEMD_SLICE":"system.slice","MESSAGE":"time=\"2019-07-01T18:53:47.371702420+02:00\" level=info msg=\"pickfirstBalancer: HandleSubConnStateChange: 0xc42002c3e0, CONNECTING\" module=grpc","_EXE":"/usr/bin/dockerd-ce","SYSLOG_IDENTIFIER":"dockerd","_SYSTEMD_UNIT":"docker.service","_SYSTEMD_CGROUP":"/system.slice/docker.service","_SYSTEMD_INVOCATION_ID":"b16ab4b640fd4721a1d56fbfc7db0941","_BOOT_ID":"6ad27b41a7994a9592b7a9f031b4748e","PRIORITY":"6","__CURSOR":"s=f9eccc9100f5452ab6e7af6c221c296b;i=3ee;b=6ad27b41a7994a9592b7a9f031b4748e;m=356270e;t=58ca178c548d6;x=31accf9bc227f784","_MACHINE_ID":"03a7024e57114636946d8fc3bf8f4863","_PID":"861","_CMDLINE":"/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock","_HOSTNAME":"jarek-VirtualBox","__REALTIME_TIMESTAMP":"1562000027371734","_SELINUX_CONTEXT":"unconfined\n","SYSLOG_FACILITY":"3","_GID":"0","__MONOTONIC_TIMESTAMP":"55977742","_STREAM_ID":"91dd1df6941e465d8f205115fe436e36","_TRANSPORT":"stdout"}
{"_SYSTEMD_INVOCATION_ID":"b16ab4b640fd4721a1d56fbfc7db0941","_MACHINE_ID":"03a7024e57114636946d8fc3bf8f4863","_PID":"861","_SYSTEMD_UNIT":"docker.service","__REALTIME_TIMESTAMP":"1562000027383560","__CURSOR":"s=f9eccc9100f5452ab6e7af6c221c296b;i=3ef;b=6ad27b41a7994a9592b7a9f031b4748e;m=3565540;t=58ca178c57708;x=a9ec1a64ca6dd3b7","_UID":"0","MESSAGE":"time=\"2019-07-01T18:53:47.383516213+02:00\" level=info msg=\"parsed scheme: \\\"unix\\\"\" module=grpc","_HOSTNAME":"jarek-VirtualBox","_GID":"0","SYSLOG_FACILITY":"3","_BOOT_ID":"6ad27b41a7994a9592b7a9f031b4748e","__MONOTONIC_TIMESTAMP":"55989568","_STREAM_ID":"91dd1df6941e465d8f205115fe436e36","_COMM":"dockerd","_CAP_EFFECTIVE":"3fffffffff","_EXE":"/usr/bin/dockerd-ce","_CMDLINE":"/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock","SYSLOG_IDENTIFIER":"dockerd","PRIORITY":"6","_SYSTEMD_CGROUP":"/system.slice/docker.service","_TRANSPORT":"stdout","_SELINUX_CONTEXT":"unconfined\n","_SYSTEMD_SLICE":"system.slice"}
{"_SYSTEMD_UNIT":"docker.service","__CURSOR":"s=f9eccc9100f5452ab6e7af6c221c296b;i=3f0;b=6ad27b41a7994a9592b7a9f031b4748e;m=356564b;t=58ca178c57813;x=bb68cea2680be4e3","__REALTIME_TIMESTAMP":"1562000027383827","_PID":"861","_HOSTNAME":"jarek-VirtualBox","PRIORITY":"6","_COMM":"dockerd","__MONOTONIC_TIMESTAMP":"55989835","_UID":"0","_SYSTEMD_SLICE":"system.slice","_CMDLINE":"/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock","SYSLOG_IDENTIFIER":"dockerd","_EXE":"/usr/bin/dockerd-ce","_MACHINE_ID":"03a7024e57114636946d8fc3bf8f4863","_SELINUX_CONTEXT":"unconfined\n","MESSAGE":"time=\"2019-07-01T18:53:47.383797349+02:00\" level=info msg=\"scheme \\\"unix\\\" not registered, fallback to default scheme\" module=grpc","SYSLOG_FACILITY":"3","_BOOT_ID":"6ad27b41a7994a9592b7a9f031b4748e","_CAP_EFFECTIVE":"3fffffffff","_STREAM_ID":"91dd1df6941e465d8f205115fe436e36","_SYSTEMD_CGROUP":"/system.slice/docker.service","_SYSTEMD_INVOCATION_ID":"b16ab4b640fd4721a1d56fbfc7db0941","_GID":"0","_TRANSPORT":"stdout"}
{"_SYSTEMD_INVOCATION_ID":"b16ab4b640fd4721a1d56fbfc7db0941","_CAP_EFFECTIVE":"3fffffffff","_BOOT_ID":"6ad27b41a7994a9592b7a9f031b4748e","PRIORITY":"6","_PID":"861","SYSLOG_FACILITY":"3","_COMM":"dockerd","__REALTIME_TIMESTAMP":"1562000027425935","_STREAM_ID":"91dd1df6941e465d8f205115fe436e36","_HOSTNAME":"jarek-VirtualBox","_SELINUX_CONTEXT":"unconfined\n","_UID":"0","_MACHINE_ID":"03a7024e57114636946d8fc3bf8f4863","__CURSOR":"s=f9eccc9100f5452ab6e7af6c221c296b;i=3f1;b=6ad27b41a7994a9592b7a9f031b4748e;m=356fac6;t=58ca178c61c8f;x=97663111e15f60a2","MESSAGE":"time=\"2019-07-01T18:53:47.425783090+02:00\" level=info msg=\"ccResolverWrapper: sending new addresses to cc: [{unix:///run/containerd/containerd.sock 0  <nil>}]\" module=grpc","__MONOTONIC_TIMESTAMP":"56031942","_SYSTEMD_CGROUP":"/system.slice/docker.service","_SYSTEMD_SLICE":"system.slice","_SYSTEMD_UNIT":"docker.service","SYSLOG_IDENTIFIER":"dockerd","_GID":"0","_TRANSPORT":"stdout","_CMDLINE":"/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock","_EXE":"/usr/bin/dockerd-ce"}`

func TestBatchIterator(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(iteratorTestData))
	batches := NewBatchIterator(3, scanner)
	allLines := make([]string, 0)
	for batches.Next() {
		lines := batches.Value()
		for _, line := range lines {
			allLines = append(allLines, line)
		}
	}
	expectedDocuments := 11
	assert.Equal(t, len(allLines), expectedDocuments)
}
