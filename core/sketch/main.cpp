#include "common.hpp"
#include "ServoController.hpp"
#include "MotorController.hpp"

int status = STATUS_OK;
ServoController servoController;
MotorController motorController =
        MotorController(EL_PIN, C1L_PIN, C2L_PIN, ER_PIN, C1R_PIN, C2R_PIN);

void receive_data(int);
void send_data();

void setup()
{
    servoController.init(SERVO_X_PIN, SERVO_Y_PIN);
    Wire.begin(ARDUINO1_ADDR);
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
        case MOVE_FORWARD:
            motorController.moveForward();
            break;
        case MOVE_BACK:
            motorController.moveBackward();
            break;
        case MOVE_LEFT:
            motorController.turnLeft();
            break;
        case MOVE_RIGHT:
            motorController.turnRight();
            break;
        case MOVE_STOP:
            motorController.stop();
            break;
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
    Wire.write("status report, captain!");
}
