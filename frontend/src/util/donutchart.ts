import { Coordinates } from "../types/donutchart";

export const getCoordinatesForPercent = (percent: number): Coordinates => {
    return [Math.cos, Math.sin].map((f) => f(2 * Math.PI * percent));
};
