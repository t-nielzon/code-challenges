function areaOfPolygonInsideCircle(circleRadius, numberOfSides) {
  const area = (numberOfSides * circleRadius * circleRadius * Math.sin(2 * Math.PI / numberOfSides)) / 2;
  return Math.round(area * 1000) / 1000;
}