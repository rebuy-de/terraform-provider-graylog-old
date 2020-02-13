resource "graylog_input" "gelf_udp" {
  title  = "GELF UDP"
  global = true

  gelf_udp {
    bind_address = "0.0.0.0"
    port         = 22201
  }
}
