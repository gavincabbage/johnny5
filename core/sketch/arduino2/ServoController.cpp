#include "ServoController.hpp"

void ServoController::init(byte xPin, byte yPin) {
    servoX.attach(xPin);
    servoY.attach(yPin);
    this->center();
}

void ServoController::lookDown() {
    this->look(servoY, false, SERVO_DOWN_MAX);
}

void ServoController::lookUp() {
    this->look(servoY, true, SERVO_UP_MAX);
}

void ServoController::lookLeft() {
    this->look(servoX, true, SERVO_LEFT_MAX);
}

void ServoController::lookRight() {
    this->look(servoX, false, SERVO_RIGHT_MAX);
}

void ServoController::center() {
    servoX.write(SERVO_CENTER);
    servoY.write(SERVO_CENTER);
}

void ServoController::look(Servo servo, bool add, int max) {
    int newDirection;
    if (add) {
        newDirection = servo.read() + SERVO_STEP;
        newDirection = newDirection > max ? max : newDirection;
    } else {
        newDirection = servo.read() - SERVO_STEP;
        newDirection = newDirection < max ? max : newDirection;
    }
    servo.write(newDirection);
}
