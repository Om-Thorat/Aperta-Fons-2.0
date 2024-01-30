import { writable } from "svelte/store";

export const state = writable("home");

export const origin = writable("http://localhost:8080")