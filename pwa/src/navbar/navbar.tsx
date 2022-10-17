import { createEffect, createSignal } from 'solid-js';
import styles from './navbar.module.css';


export const NavBar = () => {
    const [darkMode, setDarkMode] = createSignal(localStorage.getItem('theme') === 'dark' || (window.matchMedia('(prefers-color-scheme: dark)').matches));

    createEffect(() => {
        const targetTheme = darkMode() ? 'dark' : 'light';
        document.documentElement.setAttribute('data-theme', targetTheme);
        localStorage.setItem('theme', targetTheme);
    });

    return (
        <div class={styles.Navbar}>
            <h3>Sudoku</h3>
            <div class={styles.Filler} />
            <input type='checkbox' checked={darkMode()} onChange={() => setDarkMode(old => !old)} />
        </div>
    );
};