import { createSignal, createContext, useContext, JSX, Accessor } from "solid-js";

type SudokuStore = {
  board: Accessor<number[]>,
  options: Accessor<number[][]>,
  setBoardValue: (index: number, value: number) => void,
  addOption: (index: number, option: number) => void,
  removeOption: (index: number, option: number) => void,
}

const SudokuContext = createContext<SudokuStore | null>(null);

export function SudokuProvider(props: any) {
  const [board, setBoard] = createSignal<number[]>(Array.from(Array(81)).map(() => 0));
  const [options, setOptions] = createSignal<number[][]>(Array.from(Array(81)).map(() => []));
  const store: SudokuStore = {
    board,
    options,
    setBoardValue(index: number, value: number) {
      setBoard(old => [...old.slice(0, index), value, ...old.slice(index + 1)]);
    },
    addOption(index: number, option: number) {
      setOptions(old => [...old.slice(0, index), [...old[index], option], ...old.slice(index + 1)])
    },
    removeOption(index: number, option: number) {
      setOptions(old => [...old.slice(0, index), [...old[index].filter(v => v != option)], ...old.slice(index + 1)])
    }
  };

  store.setBoardValue(15, 5);
  store.setBoardValue(18, 6);
  store.setBoardValue(27, 7);
  store.addOption(65, 7);
  store.addOption(65, 3);
  store.addOption(65, 2);
  store.addOption(69, 1);
  store.addOption(69, 8);

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