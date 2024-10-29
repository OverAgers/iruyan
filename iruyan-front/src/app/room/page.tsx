"use client"
import Seats from "@/features/seats/components/seats";
import RoomLayout from "@/components/layouts/room-layout";

export default function RoomPage() {
  return (
    <RoomLayout>
      <Seats />
    </RoomLayout>
  );
}