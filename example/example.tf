provider "graylog" {
  #server_url = "http://localhost:9000"

  server_url = "http://172.18.0.4:9000/"
  username   = "admin"
  password   = "admin"
}

resource "graylog_input" "gelf_udp" {
  title  = "GELF UDP"
  type   = ""
  global = true

  gelf_udp {
    bind_address = "0.0.0.0"
    port         = 22201
  }
}
