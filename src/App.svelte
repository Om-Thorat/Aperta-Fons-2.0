<script lang="ts">
  import Hero from "./lib/Hero.svelte";
  import Timer from "./lib/Timer.svelte";

  
  let currstate:string;
  import { state } from './store';

  state.subscribe((e)=>{
    currstate = e;
    console.log(currstate)
  })
</script>

<main class="bg-black h-[100svh] flex flex-col items-center gap-4">
  {#if currstate=="home"}
    <Hero />
    <Timer final="31 dec 2023" />
  {:else if currstate=="dashboard"}
    <span>Dashboard here</span>
  {/if}
</main>
