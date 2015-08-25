#include <Arduino.h>

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

#define EL_PIN            10
#define ER_PIN            11
#define C1L_PIN           8
#define C2L_PIN           9
#define C1R_PIN           12
#define C2R_PIN           13

#define MOVE_FORWARD      10
#define MOVE_BACK         11
#define MOVE_LEFT         12
#define MOVE_RIGHT        13
#define MOVE_STOP		      14

#define PAN_CENTER        20
#define PAN_LEFT          21
#define PAN_RIGHT         22
#define PAN_UP            23
#define PAN_DOWN          24

#define STATUS_OK         42
