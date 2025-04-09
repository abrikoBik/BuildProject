import React from "react";

type DonutChartData = {
    value: number;
    color: string;
};

type DonutChartProps = {
    data: DonutChartData[];
    width: number;
    height: number;
    innerRadius: number;
    outerRadius: number;
};

const DonutChart: React.FC<DonutChartProps> = ({
    data,
    width,
    height,
    innerRadius,
    outerRadius,
}) => {
    const total = data.reduce((sum, item) => sum + item.value, 0);
    let cumulativeValue = 0;

    if (total === 0) {
        return <div>No data</div>;
    }

    const getCoordinatesForPercent = (percent: number): [number, number] => {
        const x = Math.cos(2 * Math.PI * percent);
        const y = Math.sin(2 * Math.PI * percent);
        return [x, y];
    };

    const paths = data
        .map((slice, index) => {
            if (slice.value === 0) {
                return null;
            }

            const [startX, startY] = getCoordinatesForPercent(
                cumulativeValue / total,
            );
            cumulativeValue += slice.value;
            const [endX, endY] = getCoordinatesForPercent(
                cumulativeValue / total,
            );

            const largeArcFlag = slice.value / total > 0.5 ? 1 : 0;

            let pathData: string;

            if (slice.value === total) {
                pathData = [
                    `M ${outerRadius} 0`,
                    `A ${outerRadius} ${outerRadius} 0 1 1 ${outerRadius} ${outerRadius * 2}`,
                    `A ${outerRadius} ${outerRadius} 0 1 1 ${outerRadius} 0`,
                    `M ${outerRadius} ${outerRadius - innerRadius}`,
                    `A ${innerRadius} ${innerRadius} 0 1 0 ${outerRadius} ${outerRadius + innerRadius}`,
                    `A ${innerRadius} ${innerRadius} 0 1 0 ${outerRadius} ${outerRadius - innerRadius}`,
                    "Z",
                ].join(" ");
            } else {
                pathData = [
                    `M ${outerRadius * startX + outerRadius} ${outerRadius * startY + outerRadius}`,
                    `A ${outerRadius} ${outerRadius} 0 ${largeArcFlag} 1 ${outerRadius * endX + outerRadius} ${outerRadius * endY + outerRadius}`,
                    `L ${innerRadius * endX + outerRadius} ${innerRadius * endY + outerRadius}`,
                    `A ${innerRadius} ${innerRadius} 0 ${largeArcFlag} 0 ${innerRadius * startX + outerRadius} ${innerRadius * startY + outerRadius}`,
                    "Z",
                ].join(" ");
            }

            return <path key={index} d={pathData} fill={slice.color} />;
        })
        .filter(Boolean);

    const percentage =
        total > 0 ? Math.round((data[0].value / total) * 100) : 0;

    return (
        <div
            className="relative flex items-center justify-center transition-transform"
            style={{ width: `${width}px`, height: `${height}` }}
        >
            <svg
                width={width}
                height={height}
                viewBox={`0 0 ${outerRadius * 2} ${outerRadius * 2}`}
            >
                {paths}
            </svg>
            <div className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 text-center font-semibold text-4xl">
                {`${percentage}%`}
            </div>
        </div>
    );
};

export default DonutChart;
