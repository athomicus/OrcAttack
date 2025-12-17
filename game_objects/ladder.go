components {
  id: "ladder"
  component: "/main/ladder.script"
}
embedded_components {
  id: "ladder_sprite"
  type: "sprite"
  data: "default_animation: \"ladder\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/images/orc_enemy/Orc.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.4
    y: 0.4
    z: 0.4
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"ladder\"\n"
  "mask: \"orc\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -2.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 28.619425\n"
  "  data: 71.346504\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
