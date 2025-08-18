import { Gantt, Task, ViewMode } from '@rsagiev/gantt-task-react-19';
import "@rsagiev/gantt-task-react-19/dist/index.css";

const GanttChart = () => {
    const tasks: Task[] = [
        {
            start: new Date(2020, 1, 1),
            end: new Date(2020, 1, 2),
            name: 'Idea',
            id: 'Task 0',
            type:'task',
            progress: 45,
            isDisabled: true,
            styles: { progressColor: '#ffbb54', progressSelectedColor: '#ff9e0d' },
        },
        {
            start: new Date(2020, 1, 1),
            end: new Date(2020, 6, 2),
            name: 'Idea',
            id: 'Task 1',
            type:'task',
            progress: 45,
            isDisabled: true,
            styles: { progressColor: '#ffbb54', progressSelectedColor: '#ff9e0d' },
        },
        {
            start: new Date(2020, 12, 1),
            end: new Date(2021, 1, 2),
            name: 'Idea',
            id: 'Task 2',
            type:'task',
            progress: 45,
            isDisabled: true,
            styles: { progressColor: '#ffbb54', progressSelectedColor: '#ff9e0d' },
        },
        {
            start: new Date(2020, 1, 1),
            end: new Date(2020, 1, 2),
            name: 'Idea',
            id: 'Task 3',
            type:'task',
            progress: 45,
            isDisabled: true,
            styles: { progressColor: '#ffbb54', progressSelectedColor: '#ff9e0d' },
        },
        {
            start: new Date(2020, 1, 1),
            end: new Date(2020, 1, 2),
            name: 'Idea',
            id: 'Task 4',
            type:'task',
            progress: 45,
            isDisabled: true,
            styles: { progressColor: '#ffbb54', progressSelectedColor: '#ff9e0d' },
        },
    ];

    return (
        <div className="w-full h-full">
            <Gantt
                tasks={tasks}
                viewMode={ViewMode.Week}
                locale='rus'
                listCellWidth=''
            />
        </div>
    );
}

export default GanttChart;