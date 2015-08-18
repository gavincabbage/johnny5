// arduino2/sketch.cpp
// servo subsystem

#include <Wire.h>
#include <Servo.h>

#include "../gbot.h"

Servo servo_x, servo_y;
int servo_x_direction = SERVO_CENTER;
int servo_y_direction = SERVO_CENTER;
int status = STATUS_OK;

void receive_data(int);
void send_data();
void servo_left();
void servo_right();
void servo_up();
void servo_down();
void servo_center();
void servo_move(Servo, int, int);

void setup()
{
    servo_x.attach(SERVO_X_PIN);
    servo_y.attach(SERVO_Y_PIN);
    servo_x.write(SERVO_CENTER);
    servo_y.write(SERVO_CENTER);
    Wire.begin(ARDUINO2_ADDR);
    Wire.onReceive(receive_data);
    Wire.onRequest(send_data);
}

void loop()
{
    delay(100);
}

void receive_data(int byteCount)
{
    while (Wire.available())
    {
        switch (Wire.read())
        {
        case PAN_CENTER:
            servo_center();
            break;
        case PAN_LEFT:
            servo_left();
            break;
        case PAN_RIGHT:
            servo_right();
            break;
        case PAN_UP:
            servo_up();
            break;
        case PAN_DOWN:
            servo_down();
            break;
        }
    }
}

void send_data()
{
    Wire.write(status);
}

void servo_left()
{
    servo_move(servo_x, servo_x.read() + SERVO_STEP, SERVO_LEFT_MAX);
}

void servo_right()
{
    servo_move(servo_x, servo_x.read() - SERVO_STEP, SERVO_RIGHT_MAX);
}

void servo_up()
{
    servo_move(servo_y, servo_y.read() + SERVO_STEP, SERVO_UP_MAX);
}

void servo_down()
{
    servo_move(servo_y, servo_y.read() - SERVO_STEP, SERVO_DOWN_MAX);
}

void servo_move(Servo servo, int newDirection, int max)
{
    newDirection = newDirection > max ? max : newDirection;
    servo.write(newDirection);
}

void servo_center()
{
    servo_x.write(SERVO_CENTER);
    servo_y.write(SERVO_CENTER);
}
