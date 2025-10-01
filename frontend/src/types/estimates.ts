export interface LineItem {
    id: number,
    name: string,
    unit: string,
    count: number,
    cost: number,
};

export interface EditableCellProps {
    value: string;
    onSave?: (newValue: string) => void;
    onChange?: (newValue: string) => void;
}

export interface EstimateGroup {
    id: number
    name: string;
}