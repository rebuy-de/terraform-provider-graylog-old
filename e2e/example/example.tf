resource "graylog_input" "gelf_tcp" {
  title  = "GELF TCP"
  global = true

  gelf_tcp {
    bind_address = "0.0.0.0"
    port         = 12201
  }
}

resource "graylog_input" "gelf_udp" {
  title  = "GELF UDP"
  global = true

  gelf_udp {
    bind_address = "0.0.0.0"
    port         = 22201
  }
}

resource "graylog_input" "beats" {
  title  = "Beats"
  global = true

  beats {
    bind_address = "0.0.0.0"
    port         = 5044
  }
}

resource "graylog_input" "gelf_http" {
  title  = "GELF HTTP"
  global = true

  gelf_http {
    bind_address = "0.0.0.0"
    port         = 12202
  }
}
