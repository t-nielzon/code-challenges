function roomMates(rooms, floor) {
  const start = (floor - 1) * 6;
  return rooms.filter((name, index) => index >= start && index < start + 6 && name !== "");
}

module.exports = { roomMates };