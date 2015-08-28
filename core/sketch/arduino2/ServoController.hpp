#ifndef ServoController_H
#define ServoController_H

#include "../common.hpp"
#include <Servo.h>

class ServoController {

    public:
        void init(byte xPin, byte yPin);
        void lookDown();
        void lookUp();
        void lookLeft();
        void lookRight();
        void center();

    private:
        Servo servoX;
        Servo servoY;
        void look(Servo, bool, int);
};

#endif
