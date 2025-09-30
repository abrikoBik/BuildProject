export interface LineItem {
    name: string,
    unit: string,
    amount: number,
    price: number,
    cost: number
};

export interface EditableCellProps {
    value: string;
    onSave?: (newValue: string) => void;
    onChange?: (newValue: string) => void;
}

export interface EstimateGroup {
    name: string;
}