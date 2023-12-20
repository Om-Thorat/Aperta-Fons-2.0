<script lang="ts">
    let pars: Array<any>;
    let vis = false;
    async function getpars(){
        let s = await fetch("http://localhost:8080/all");
        pars = await s.json();
        console.log(pars)
        vis = true
    }
    getpars()
</script>


<div class="pt-12 flex flex-col w-3/5">
    <div class="flex items-center">
        <span class="text-6xl w-[70%] font-mono font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-[#12c2e9] via-[#c471ed] to-[#f64f59]">Dashboard</span>
        <select class="select h-4/5 text-lg font-mono w-[30%] bg-transparent rounded-3xl focus:outline-none select-primary max-w-xs">
            <option>Your PRS</option>
            <option selected>Leaderboard</option>
        </select>          
</div>
<div class="divider divider-info"></div> 
<div class="flex">
    <span class="text-2xl w-3/5 text-center">Pr Title</span>
    <div class="divider divider-horizontal divider-info"></div>
    <span class="text-2xl w-2/5 text-center">Time Added</span>
</div>

{#if vis}
{#each pars as { Name,Prs }}
<div class="flex">
        <span class="text-2xl w-3/5 text-center">{Name}</span>
        <div class="divider divider-horizontal divider-info"></div>
        <span class="text-2xl w-2/5 text-center">{Prs}</span>
</div>
{/each}
{/if}
</div>
