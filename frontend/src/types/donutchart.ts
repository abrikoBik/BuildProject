export type DonutChartData = {
    value: number;
    color: string;
};

export type DonutChartProps = {
    data: DonutChartData[];
    width: number;
    height: number;
    innerRadius: number;
    outerRadius: number;
    animationDuration?: number;
};

export type Coordinates = number[];
