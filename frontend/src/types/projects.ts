export type ValuePiece = Date | null;

export type Value = ValuePiece | [ValuePiece, ValuePiece];

export interface ProjectRequest  {
    name: string,
    userID: number,
    description: string,
    startTime: number,
    endTime: number,
    latitude: number,
    longitude: number,
    id: number
}