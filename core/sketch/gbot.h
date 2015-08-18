#include <Arduino.h>

#define LM_PIN_1          2      // IN1, driver 1
#define LM_PIN_2          3      // IN2, driver 1
#define LM_PIN_3          4      // IN3, driver 1
#define LM_PIN_4          5      // IN4, driver 1
#define RM_PIN_1          6      // IN1, driver 2
#define RM_PIN_2          7      // IN2, driver 2
#define RM_PIN_3          8      // IN3, driver 2
#define RM_PIN_4          9      // IN4, driver 2
#define SERVO_X_PIN       10
#define SERVO_Y_PIN       11

#define ARDUINO1_ADDR     0x04
#define ARDUINO2_ADDR     0x05

#define SERVO_LEFT_MAX    160
#define SERVO_RIGHT_MAX   20
#define SERVO_CENTER      90
#define SERVO_UP_MAX      160
#define SERVO_DOWN_MAX    70
#define SERVO_STEP        10
#define SERVO_INTERVAL    5000

#define MAX_SPEED         1000.0
#define SPD_LEFT_FORWARD  -750.0
#define SPD_LEFT_BACK     750.0
#define SPD_LEFT_LEFT	  400.0
#define SPD_LEFT_RIGHT    -400.0
#define SPD_RIGHT_FORWARD 750.0
#define SPD_RIGHT_BACK    -750.0
#define SPD_RIGHT_LEFT	  400.0
#define SPD_RIGHT_RIGHT   -400.0
#define SPD_STOP		  0.0

#define MOVE_FORWARD      10
#define MOVE_BACK         11
#define MOVE_LEFT         12
#define MOVE_RIGHT        13
#define MOVE_STOP		  14

#define PAN_CENTER        20
#define PAN_LEFT          21
#define PAN_RIGHT         22
#define PAN_UP            23
#define PAN_DOWN          24

#define STATUS_OK         42
