import { SeatsInfo } from './seats-info';

export type RoomInfo = {
  roomId: string;
  roomName: string;
  seats: SeatsInfo[];
};