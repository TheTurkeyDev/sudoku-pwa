export type BoardType = {
    readonly board: readonly (readonly number[])[]
    readonly solution: readonly (readonly number[])[]
    readonly options: readonly (readonly (readonly number[])[])[]
}