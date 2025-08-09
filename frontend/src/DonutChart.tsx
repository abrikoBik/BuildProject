import React, { useState, useEffect, useRef } from "react";
import { DonutChartProps } from "./types/donutchart";
import { getCoordinatesForPercent } from "./util/donutchart";

const DonutChart: React.FC<DonutChartProps> = ({
  data,
  width,
  height,
  innerRadius,
  outerRadius,
  animationDuration = 1000,
}) => {
  const total = data.reduce((sum, item) => sum + item.value, 0);
  const [displayData, setDisplayData] = useState(data);
  const animationRef = useRef<number>(0);
  const startTimeRef = useRef<number>(0);
  const startValuesRef = useRef<number[]>(data.map(item => item.value));

  if (total === 0) {
    return <div>No data</div>;
  }

  useEffect(() => {
    // Сохраняем текущие отображаемые значения как начальные для анимации
    startValuesRef.current = displayData.map(item => item.value);
    startTimeRef.current = 0;
    
    const animate = (timestamp: number) => {
      if (!startTimeRef.current) {
        startTimeRef.current = timestamp;
      }

      const elapsed = timestamp - startTimeRef.current;
      const progress = Math.min(elapsed / animationDuration, 1);
      
      // Интерполяция между текущими отображаемыми и новыми значениями
      const newDisplayData = data.map((item, i) => {
        const startValue = startValuesRef.current[i] || 0;
        const targetValue = item.value;
        return {
          ...item,
          value: startValue + (targetValue - startValue) * progress
        };
      });

      setDisplayData(newDisplayData);

      if (progress < 1) {
        animationRef.current = requestAnimationFrame(animate);
      }
    };

    animationRef.current = requestAnimationFrame(animate);

    return () => {
      if (animationRef.current) {
        cancelAnimationFrame(animationRef.current);
      }
    };
  }, [data, animationDuration]);

  let cumulativeValue = 0;
  const paths = displayData.map((slice, index) => {
    if (slice.value === 0) return null;

    const slicePercent = slice.value / total;
    const [startX, startY] = getCoordinatesForPercent(cumulativeValue);
    cumulativeValue += slicePercent;
    const [endX, endY] = getCoordinatesForPercent(cumulativeValue);

    const largeArcFlag = slicePercent > 0.5 ? 1 : 0;
    let pathData: string;

    if (Math.abs(slicePercent - 1) < 0.0001) {
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

    return (
      <path
        key={index}
        d={pathData}
        fill={slice.color}
      />
    );
  });

  const percentage = total > 0 ? Math.round((data[0].value / total) * 100) : 0;

  return (
    <div
      className="relative flex items-center justify-center"
      style={{ width: `${width}px`, height: `${height}px` }}
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