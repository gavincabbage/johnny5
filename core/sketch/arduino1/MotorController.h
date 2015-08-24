#ifndef MotorController_H
#define MotorController_H

#include <Arduino.h>

#define DELAY 10
#define C1L_FORWARD HIGH
#define C2L_FORWARD LOW
#define C1R_FORWARD LOW
#define C2R_FORWARD HIGH
#define C1L_BACKWARD LOW
#define C2L_BACKWARD HIGH
#define C1R_BACKWARD HIGH
#define C2R_BACKWARD LOW

class MotorController {

    public:
        MotorController(byte eL, byte c1L, byte c2L, byte eR, byte c1R, byte c2R);
        void moveForward();
        void moveBackward();
        void turnLeft();
        void turnRight();
        void stop();

    private:
        byte enableLeft;
        byte control1Left;
        byte control2Left;
        byte enableRight;
        byte control1Right;
        byte control2Right;
        void preamble();
};

#endif
