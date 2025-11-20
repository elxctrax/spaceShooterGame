components {
  id: "script"
  component: "/stars/factory.script"
}
embedded_components {
  id: "star_factory"
  type: "factory"
  data: "prototype: \"/stars/star.go\"\n"
  ""
}
embedded_components {
  id: "bonus_factory"
  type: "factory"
  data: "prototype: \"/stars/bonus_star.go\"\n"
  ""
}
embedded_components {
  id: "bad_factory"
  type: "factory"
  data: "prototype: \"/stars/bad_star.go\"\n"
  ""
}
