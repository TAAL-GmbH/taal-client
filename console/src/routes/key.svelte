<script>
    export let key;
</script>

<article>
    <small>
      api key: <b>{key.api_key}</b>   
    </small><br/>
    <small>
      address: <b>{key.address}</b>   
    </small><br/>
    <small>
      public key: <b>{key.public_key}</b>
    </small><br/>
</article>

<style>
    article {
      margin: 0 0 1em 0;
    }
    h1 {
      font-size: 1.4em;
      margin: 0;
      display: block;
    }
</style>
