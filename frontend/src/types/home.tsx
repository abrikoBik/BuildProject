export interface TaskGroup {
    id: number;
    name: string | null;
}

export interface TaskRequest {
    startTime: number;
    endTime: number;
    name: string;
}