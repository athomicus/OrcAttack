
---
# OrcAttack

stateDiagram-v2
    direction LR

    [*] --> IDLE : start

    IDLE --> WALKING : click on empty ground
    IDLE --> WALKING_FOR_STONE : click on stone

    WALKING --> IDLE : reached destination
    WALKING --> WALKING_FOR_STONE : click on stone

    WALKING_FOR_STONE --> HOLDING_STONE : arrived at stone, grab it
    HOLDING_STONE --> THROWING_STONE : click to throw
    THROWING_STONE --> IDLE : throw finished

    WALKING_FOR_STONE --> IDLE : cancelled pickup