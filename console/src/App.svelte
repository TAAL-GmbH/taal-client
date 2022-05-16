<script>
  // import { onMount } from "svelte";

  // let name;

  // onMount(async () => {
  //   name = "N/A";

  //   const response = await fetch(`${BASE_URL}/api/v1/test`);
  //   const data = await response.text();

  //   name = data;
  // });

  import { Router, Route, Link } from "svelte-routing";
  import Home from "./routes/home.svelte";
  import Settings from "./routes/settings.svelte";

  export let url = "";
</script>

<Router {url}>
  <nav>
    <Link to="/">Home</Link>
    <Link to="/	settings">Settings</Link>
  </nav>

  <div>
    <Route path="/"><Home /></Route>
    <Route path="settings" component={Settings} />
  </div>
</Router>
