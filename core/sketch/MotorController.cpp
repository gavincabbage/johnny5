#include "MotorController.hpp"

#define DELAY 10
#define C1L_FORWARD HIGH
#define C2L_FORWARD LOW
#define C1R_FORWARD LOW
#define C2R_FORWARD HIGH
#define C1L_BACKWARD LOW
#define C2L_BACKWARD HIGH
#define C1R_BACKWARD HIGH
#define C2R_BACKWARD LOW

MotorController::MotorController(byte eL, byte c1L, byte c2L, byte eR, byte c1R, byte c2R) {
    enableLeft = eL;
    control1Left = c1L;
    control2Left = c2L;
    enableRight = eR;
    control1Right = c1R;
    control2Right = c2R;
    speed = 255;

    pinMode(enableLeft, OUTPUT);
    pinMode(control1Left, OUTPUT);
    pinMode(control2Left, OUTPUT);
    pinMode(enableRight, OUTPUT);
    pinMode(control1Right, OUTPUT);
    pinMode(control2Right, OUTPUT);
}

void MotorController::moveForward() {
    preamble();
    digitalWrite(control1Left, C1L_FORWARD);
    digitalWrite(control2Left, C2L_FORWARD);
    digitalWrite(control1Right, C1R_FORWARD);
    digitalWrite(control2Right, C2R_FORWARD);
}

void MotorController::moveBackward() {
    preamble();
    digitalWrite(control1Left, C1L_BACKWARD);
    digitalWrite(control2Left, C2L_BACKWARD);
    digitalWrite(control1Right, C1R_BACKWARD);
    digitalWrite(control2Right, C2R_BACKWARD);
}

void MotorController::turnLeft() {
    preamble();
    digitalWrite(control1Left, C1L_BACKWARD);
    digitalWrite(control2Left, C2L_BACKWARD);
    digitalWrite(control1Right, C1R_FORWARD);
    digitalWrite(control2Right, C2R_FORWARD);
}

void MotorController::turnRight() {
    preamble();
    digitalWrite(control1Left, C1L_FORWARD);
    digitalWrite(control2Left, C2L_FORWARD);
    digitalWrite(control1Right, C1R_BACKWARD);
    digitalWrite(control2Right, C2R_BACKWARD);
}

void MotorController::stop() {
    digitalWrite(enableLeft, LOW);
    digitalWrite(enableRight, LOW);
}

void MotorController::setSpeed(int newSpeed) {
    speed = newSpeed;
}

int MotorController::getSpeed() {
    return speed;
}

void MotorController::preamble() {
    digitalWrite(enableLeft, LOW);
    digitalWrite(enableRight, LOW);
    delay(DELAY);
    digitalWrite(enableLeft, HIGH);
    digitalWrite(enableRight, HIGH);
    // analogWrite(enableLeft, speed);
    // analogWrite(enableRight, speed);
}
