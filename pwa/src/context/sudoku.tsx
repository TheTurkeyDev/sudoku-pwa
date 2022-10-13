import { createSignal, createContext, useContext, Setter, Accessor } from "solid-js";
import { BoardType } from "../board/board-type";

type SudokuStore = {
  board: Accessor<number[]>,
  options: Accessor<number[][]>,
  loadBoard: (board: BoardType) => void,
  setBoardValue: (index: number, value: number) => void,
  addOption: (index: number, option: number) => void,
  removeOption: (index: number, option: number) => void,
  toggleOption: (index: number, option: number) => void,
  onInput: (value: number) => void,
  setSelectedCell: Setter<number>,
  selectedCell: Accessor<number>,
  setEditingOptions: Setter<boolean>,
  editingOptions: Accessor<boolean>,
  isLockedValue: (index: number) => boolean,
}

const SudokuContext = createContext<SudokuStore | null>(null);

export function SudokuProvider(props: any) {
  const [baseBoard, setBaseBoard] = createSignal<number[]>(Array.from(Array(81)).map(() => 0));
  const [board, setBoard] = createSignal<number[]>(Array.from(Array(81)).map(() => 0));
  const [options, setOptions] = createSignal<number[][]>(Array.from(Array(81)).map(() => []));
  const [selectedCell, setSelectedCell] = createSignal<number>(-1);
  const [editingOptions, setEditingOptions] = createSignal<boolean>(false);

  const loadBoard = (board: BoardType) => {
    const flat = board.board.flatMap(n => n)
    setBoard(flat)
    setBaseBoard(flat)
    setOptions(board.options.flatMap(o => o))
    setSelectedCell(-1)
    setEditingOptions(false)
  }

  const setBoardValue = (index: number, value: number) => {
    if (index < 0 || index > 80 || baseBoard()[index] !== 0)
      return;
    setBoard(old => [...old.slice(0, index), value, ...old.slice(index + 1)]);
  }

  const addOption = (index: number, option: number) => {
    if (index < 0 || index > 80)
      return;
    setOptions(old => [...old.slice(0, index), [...old[index], option], ...old.slice(index + 1)])
  }

  const removeOption = (index: number, option: number) => {
    if (index < 0 || index > 80)
      return;
    setOptions(old => [...old.slice(0, index), [...old[index].filter(v => v != option)], ...old.slice(index + 1)])
  }

  const toggleOption = (index: number, option: number) => {
    options()[index].includes(option) ? removeOption(index, option) : addOption(index, option)
  }

  const onInput = (value: number) => {
    if (editingOptions())
      toggleOption(selectedCell(), value)
    else
      setBoardValue(selectedCell(), value);
  }

  const isLockedValue = (index: number) => (baseBoard()[index] !== 0)

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
    isLockedValue(index: number) { return baseBoard()[index] !== 0 },
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