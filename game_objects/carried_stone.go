components {
  id: "sprite"
  component: "/assets/stone.sprite"
  position {
    x: 3.0
    y: -6.0
  }
  scale {
    x: 0.2
    y: 0.2
    z: 0.2
  }
}
components {
  id: "stone"
  component: "/main/stone.script"
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"stone\"\n"
  "mask: \"orc\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -1.0\n"
  "      y: 1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 6.4863114\n"
  "  data: 6.7697787\n"
  "  data: 10.4\n"
  "}\n"
  ""
}
