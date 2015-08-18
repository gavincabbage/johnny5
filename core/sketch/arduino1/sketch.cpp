// arduino1/sketch.cpp
// stepper motor subsystem

#include <Wire.h>
#include <AccelStepper.h>

#include "../gbot.h"

AccelStepper left_motor(AccelStepper::HALF4WIRE, LM_PIN_1, LM_PIN_3, LM_PIN_2, LM_PIN_4);
AccelStepper right_motor(AccelStepper::HALF4WIRE, RM_PIN_1, RM_PIN_3, RM_PIN_2, RM_PIN_4);
int motor_mode = MOVE_STOP;
int status = STATUS_OK;

void move_forward();
void move_back();
void move_left();
void move_right();
void move_stop();
void move(int, int, int);
void receive_data(int);
void send_data();
void setup_motors();

void setup()
{
    setup_motors();
    Wire.begin(ARDUINO1_ADDR);
    Wire.onReceive(receive_data);
    Wire.onRequest(send_data);
}

void loop()
{
    left_motor.runSpeed();
    right_motor.runSpeed();
}

void receive_data(int byteCount)
{
    while (Wire.available())
    {
        switch (Wire.read())
        {
        case MOVE_FORWARD:
            move_forward();
            break;
        case MOVE_BACK:
            move_back();
            break;
        case MOVE_LEFT:
            move_left();
            break;
        case MOVE_RIGHT:
            move_right();
            break;
        case MOVE_STOP:
            move_stop();
            break;
        }
    }
}

void move_forward()
{
    move(SPD_LEFT_FORWARD, SPD_RIGHT_FORWARD, MOVE_FORWARD);
}

void move_back()
{
    move(SPD_LEFT_BACK, SPD_RIGHT_BACK, MOVE_BACK);
}

void move_left()
{
    move(SPD_LEFT_LEFT, SPD_RIGHT_LEFT, MOVE_LEFT);
}

void move_right()
{
    move(SPD_LEFT_RIGHT, SPD_RIGHT_RIGHT, MOVE_RIGHT);
}

void move_stop()
{
    move(SPD_STOP, SPD_STOP, MOVE_STOP);
}

void move(int left_direction, int right_direction, int mode)
{
    left_motor.setSpeed(left_direction);
    right_motor.setSpeed(right_direction);
    motor_mode = mode;
}

void send_data()
{
    Wire.write(status);
}

void setup_motors()
{
    left_motor.setMaxSpeed(MAX_SPEED);
    left_motor.setSpeed(SPD_STOP);
    right_motor.setMaxSpeed(MAX_SPEED);
    right_motor.setSpeed(SPD_STOP);
}