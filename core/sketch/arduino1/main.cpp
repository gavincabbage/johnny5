#include <Wire.h>

#include "../gbot.h"
#include "MotorController.h"

int status = STATUS_OK;
MotorController motorController =
    MotorController(EL_PIN, C1L_PIN, C2L_PIN, ER_PIN, C1R_PIN, C2R_PIN);

void receive_data(int);
void send_data();

void setup()
{
  Wire.begin(ARDUINO1_ADDR);
  Wire.onReceive(receive_data);
  Wire.onRequest(send_data);
}

void loop()
{
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
    }
  }
}

void send_data()
{
  Wire.write(status);
}
