<script lang="ts">
  import Dash from "./lib/Dash.svelte";
  import Hero from "./lib/Hero.svelte";
  import Timer from "./lib/Timer.svelte";
  import Login from "./lib/Login.svelte"
  import { onMount } from "svelte";

  let currstate:string;
  import { state,origin } from './store';

  state.subscribe((e)=>{
    currstate = e;
    console.log(currstate)
  })

  onMount(()=>{
    if (window.location.href.includes("localhost")){
      // do nothing;
    } else {
      origin.set(window.location.href)
    };
    localStorage.getItem("user") ? console.log("loggedin") : document.getElementById("login").showModal();
  })
</script>

<main class="bg-black h-[100svh] flex flex-col items-center gap-4">
  {#if currstate=="home"}
    <Hero />
    <Timer final="31 dec 2023" />
    <Login />
  {:else if currstate=="dashboard"}
    <Dash />
  {/if}
</main>
