<script>
  import { onMount } from "svelte";

    // @ts-nocheck
    import { origin } from "../store";
    let pars;
    let vis = false;
    let user;
    async function getpars(){
        let s = await fetch(`${origin}/all`);
        pars = await s.json();
        console.log(pars)
        vis = true
    }
    async function syncprs() {
        let l = await fetch(`${origin}/getuser/${user}`)
        l = await l.json();
        let m = l.Prs[l.Prs.length-1]
        console.log(m);
        let prs = await fetch("https://api.github.com/search/issues?q=state%3Aclosed+author%3A${user}+type%3Apr")
        prs = await prs.json();
        console.log(prs);
        let prl = prs.items;
        console.log(prl);
        for(let i = 0;i<prl.length;i++){
          let n = new Date(prl[i].created_at).getTime()
          let t = new Date("07 Dec 23").getTime()
          if (prl[i].id == m.ID){
            break;
          }
          if(n>t){
            console.log("nice");
            console.log(prl[i]);
            fetch(`${origin}/inpr`, {
                method: "POST",
                body: JSON.stringify({
                    Name: prl[i].user.login.toLowerCase(),
                    Title: prl[i].title,
                    Link: prl[i].html_url,
                    ID: `${prl[i].id}`
            }),
            headers: {
                "Content-type": "application/json; charset=UTF-8"
            }
            });
          }
        }
    }
    getpars()
    onMount(()=>{user=localStorage.getItem("user")})
</script>


<div class="pt-12 flex flex-col w-3/5">
    <div class="flex items-center gap-5">
        <span class="text-6xl w-[70%] font-mono font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-[#12c2e9] via-[#c471ed] to-[#f64f59]">Dashboard</span>
        <button on:click={syncprs} class="btn btn-outline rounded-2xl"> Sync Prs </button>
        <select class="select h-4/5 text-lg font-mono w-[30%] bg-transparent rounded-3xl focus:outline-none select-primary max-w-xs">
            <option>Your PRS</option>
            <option selected>Leaderboard</option>
        </select>          
</div>
<div class="divider divider-info"></div> 
<div class="flex">
    <span class="text-2xl w-1/5 text-center">Rank</span>
    <div class="divider divider-horizontal divider-info"></div>    
    <span class="text-2xl w-3/5 text-center">Name</span>
    <div class="divider divider-horizontal divider-info"></div>
    <span class="text-2xl w-1/5 text-center">Total PRs</span>
</div>

<div class="flex flex-col gap-6 pt-6">
{#if vis}
{#each pars as { Name,Prs },i (i)}
<div class="flex">
    <span class="text-2xl w-1/5 text-center">{i+1}</span>
    <!-- <div class="divider divider-horizontal divider-info"></div>     -->
    <span class="text-2xl w-3/5 text-center">{Name}</span>
    <!-- <div class="divider divider-horizontal divider-info"></div> -->
    <span class="text-2xl w-1/5 text-center">{Prs ? Prs.length : 0}</span>
</div>
{/each}
{/if}
</div>
</div>