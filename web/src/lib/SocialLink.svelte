<script>
    export let icon;
    export let name;
    export let href;
</script>

<a href={href} target="_blank" rel="noreferrer">
    <img src={icon} alt={name}>
    {name}
</a>