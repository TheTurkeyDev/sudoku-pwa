import { createSignal, createContext, useContext, Setter, Accessor, JSX } from 'solid-js';
import { BoardType } from '../board/board-type';

type SudokuStore = {
    readonly board: Accessor<readonly number[]>,
    readonly options: Accessor<readonly (readonly number[])[]>,
    readonly loadBoard: (board: BoardType) => void,
    readonly setBoardValue: (index: number, value: number) => void,
    readonly addOption: (index: number, option: number) => void,
    readonly removeOption: (index: number, option: number) => void,
    readonly toggleOption: (index: number, option: number) => void,
    readonly onInput: (value: number) => void,
    readonly setSelectedCell: Setter<number>,
    readonly selectedCell: Accessor<number>,
    readonly setEditingOptions: Setter<boolean>,
    readonly editingOptions: Accessor<boolean>,
    readonly isLockedValue: (index: number) => boolean,
}

const SudokuContext = createContext<SudokuStore | null>(null);

type SudokuProviderProps = {
    readonly children: JSX.Element
};

export function SudokuProvider({ children }: SudokuProviderProps) {
    const [baseBoard, setBaseBoard] = createSignal<readonly number[]>(Array.from(Array(81)).map(() => 0));
    const [board, setBoard] = createSignal<readonly number[]>(Array.from(Array(81)).map(() => 0));
    const [solution, setSolution] = createSignal<readonly number[]>(Array.from(Array(81)).map(() => 0));
    const [options, setOptions] = createSignal<readonly (readonly number[])[]>(Array.from(Array(81)).map(() => []));
    const [selectedCell, setSelectedCell] = createSignal<number>(-1);
    const [editingOptions, setEditingOptions] = createSignal<boolean>(false);

    const loadBoard = (board: BoardType) => {
        const flat = board.board.flatMap(n => n);
        setBoard(flat);
        setBaseBoard(flat);
        setSolution(board.solution.flatMap(n => n));
        setOptions(board.options.flatMap(o => o));
        setSelectedCell(-1);
        setEditingOptions(false);
    };

    const setBoardValue = (index: number, value: number) => {
        if (index < 0 || index > 80 || baseBoard()[index] !== 0)
            return;
        setBoard(old => [...old.slice(0, index), value, ...old.slice(index + 1)]);
    };

    const addOption = (index: number, option: number) => {
        if (index < 0 || index > 80)
            return;
        setOptions(old => [...old.slice(0, index), [...old[index], option], ...old.slice(index + 1)]);
    };

    const removeOption = (index: number, option: number) => {
        if (index < 0 || index > 80)
            return;
        setOptions(old => [...old.slice(0, index), [...old[index].filter(v => v !== option)], ...old.slice(index + 1)]);
    };

    const toggleOption = (index: number, option: number) => {
        options()[index].includes(option) ? removeOption(index, option) : addOption(index, option);
    };

    const onInput = (value: number) => {
        if (editingOptions())
            toggleOption(selectedCell(), value);
        else
            setBoardValue(selectedCell(), value);
    };

    const isLockedValue = (index: number) => (baseBoard()[index] !== 0);

    const store: SudokuStore = {
        board,
        options,
        loadBoard,
        setBoardValue,
        addOption,
        removeOption,
        toggleOption,
        onInput,
        setSelectedCell,
        selectedCell,
        setEditingOptions,
        editingOptions,
        isLockedValue(index: number) { return baseBoard()[index] !== 0; },
    };

    return (
        <SudokuContext.Provider value={store}>
            {children}
        </SudokuContext.Provider>
    );
}

export const useSudoku = () => {
    const sudoku = useContext(SudokuContext);
    if (!sudoku)
        throw new Error('Sudoku is undefined! Must be used from within a Sudoku Provider!');
    return sudoku;
};