import { createSignal, createContext, useContext, Setter, Accessor, JSX, createEffect, JSXElement } from 'solid-js';
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
    readonly difficulty: Accessor<number>,
    readonly setDifficulty: Setter<number>,
    readonly startTime: Accessor<number>,
    readonly endTime: Accessor<number>,
    readonly solution: Accessor<readonly number[]>,
    readonly gameState: Accessor<number>,
    readonly setGameState: Setter<number>,
}

const SudokuContext = createContext<SudokuStore | null>(null);

type SudokuProviderProps = {
    readonly children: JSXElement
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function SudokuProvider(props: SudokuProviderProps) {
    const [gameState, setGameState] = createSignal<number>(0);
    const [baseBoard, setBaseBoard] = createSignal<readonly number[]>(Array.from(Array(81)).map(() => 0));
    const [board, setBoard] = createSignal<readonly number[]>(Array.from(Array(81)).map(() => 0));
    const [solution, setSolution] = createSignal<readonly number[]>(Array.from(Array(81)).map(() => 0));
    const [options, setOptions] = createSignal<readonly (readonly number[])[]>(Array.from(Array(81)).map(() => []));
    const [selectedCell, setSelectedCell] = createSignal<number>(-1);
    const [editingOptions, setEditingOptions] = createSignal<boolean>(false);
    const [difficulty, setDifficulty] = createSignal<number>(-1);
    const [startTime, setStartTime] = createSignal<number>(0);
    const [endTime, setEndTime] = createSignal<number>(0);

    createEffect(() => {
        if (gameState() === 1 && board().find((v, i) => !v || solution()[i] !== v) === undefined) {
            setGameState(2);
            setEndTime((new Date()).getTime());
        }
    });

    const loadBoard = (board: BoardType) => {
        const flat = board.board.flatMap(n => n);
        setBoard(flat);
        setBaseBoard(flat);
        setSolution(board.solution.flatMap(n => n));
        setOptions(board.options.flatMap(o => o));
        setSelectedCell(-1);
        setEditingOptions(false);
        setStartTime((new Date()).getTime());
        setGameState(1);

        //TODO: REMOVE!
        setTimeout(() => {
            setBoard(solution());
        }, 5000);
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
        difficulty,
        setDifficulty,
        startTime,
        endTime,
        solution,
        gameState,
        setGameState
    };

    return (
        <SudokuContext.Provider value={store}>
            {props.children}
        </SudokuContext.Provider>
    );
}

export const useSudoku = () => {
    const sudoku = useContext(SudokuContext);
    if (!sudoku)
        throw new Error('Sudoku is undefined! Must be used from within a Sudoku Provider!');
    return sudoku;
};