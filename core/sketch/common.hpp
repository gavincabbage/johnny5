#include <Arduino.h>
#include <Wire.h>

#define SERVO_X_PIN       10
#define SERVO_Y_PIN       11

#define ARDUINO1_ADDR     4
#define ARDUINO2_ADDR     5

#define SERVO_LEFT_MAX    180
#define SERVO_RIGHT_MAX   0
#define SERVO_CENTER      90
#define SERVO_UP_MAX      180
#define SERVO_DOWN_MAX    70
#define SERVO_STEP        10

#define EL_PIN            9
#define ER_PIN            10
#define C1L_PIN           7
#define C2L_PIN           8
#define C1R_PIN           11
#define C2R_PIN           12

#define MOVE_FORWARD      10
#define MOVE_BACK         11
#define MOVE_LEFT         12
#define MOVE_RIGHT        13
#define MOVE_STOP		  14

#define LOOK_CENTER        20
#define LOOK_LEFT          21
#define LOOK_RIGHT         22
#define LOOK_UP            23
#define LOOK_DOWN          24

#define STATUS_OK         42
