import { writable } from "svelte/store";

export const state = writable("home");

export const origin = "http://localhost:8080"