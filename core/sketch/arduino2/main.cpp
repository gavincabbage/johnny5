#include "../common.hpp"
#include "ServoController.hpp"

int status = STATUS_OK;
ServoController servoController;

void receive_data(int);
void send_data();

void setup()
{
    servoController.init(SERVO_X_PIN, SERVO_Y_PIN);
    Wire.begin(ARDUINO2_ADDR);
    Wire.onReceive(receive_data);
    Wire.onRequest(send_data);
}

void loop()
{
    delay(10);
}

void receive_data(int byteCount)
{
    while (Wire.available())
    {
        switch (Wire.read())
        {
        case LOOK_CENTER:
            servoController.center();
            break;
        case LOOK_LEFT:
            servoController.lookLeft();
            break;
        case LOOK_RIGHT:
            servoController.lookRight();
            break;
        case LOOK_UP:
            servoController.lookUp();
            break;
        case LOOK_DOWN:
            servoController.lookDown();
            break;
        }
    }
}

void send_data()
{
    Wire.write(status);
}
