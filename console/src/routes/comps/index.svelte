<script>
  import Button from '../../lib/components/button/index.svelte'
  import Dropdown from '../../lib/components/dropdown/index.svelte'
  import TextInput from '../../lib/components/text-input/index.svelte'
  import Checkbox from '../../lib/components/checkbox/index.svelte'
  import Radio from '../../lib/components/radio/index.svelte'
  import Progress from '../../lib/components/progress/index.svelte'
  import Switch from '../../lib/components/switch/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageWithMenu from '../../lib/components/page/template/menu/index.svelte'
  import Row from '../../lib/components/layout/row/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'

  function onChange(e) {
    console.log('onChange: detail = ', e.detail)
  }
</script>

<PageWithMenu>
  <div class="island">
    <Heading value="Components demo" />
    <Spacer h={24} />
    <div class="sub-row">
      <Heading value="Add something" size={2} />
      <Button icon="plus">Add new</Button>
    </div>
    <Spacer h={24} />
    <Dropdown
      name="bob"
      label="I am a dropdown"
      items={[
        { value: 0, label: 'Option 0' },
        { value: 1, label: 'Option 1' },
        { value: 99, label: 'Option 2' },
      ]}
      on:change={onChange}
    />
    <Spacer h={24} />
    <TextInput name="sally" label="I am a textinput" on:change={onChange} />
    <Spacer h={24} />
    <Row gap={30}>
      <Switch name="switch1" label="Switch" on:change={onChange} />
      <Switch
        name="switch2"
        label="Switch"
        labelPlacement="right"
        on:change={onChange}
      />
      <Switch name="switch3" on:change={onChange} />
    </Row>
    <Spacer h={24} />
    <Row gap={30}>
      <Radio name="radio1" label="Radio" on:change={onChange} />
      <Radio
        name="radio1"
        label="Radio"
        labelPlacement="right"
        on:change={onChange}
      />
      <Radio name="radio2" on:change={onChange} />
    </Row>
    <Spacer h={24} />
    <Row gap={30}>
      <Checkbox name="checkbox1" label="Checkbox" on:change={onChange} />
      <Checkbox
        name="checkbox2"
        label="Checkbox"
        labelPlacement="right"
        on:change={onChange}
      />
      <Checkbox name="checkbox3" on:change={onChange} />
    </Row>
    <Spacer h={24} />
    <Progress value={55} total={100} warnValue={50} dangerValue={20} />
    <!-- <Row flex>
      <div class="block" />
      <div class="block" />
      <div class="block" />
    </Row> -->
  </div>
</PageWithMenu>

<style>
  .island {
    display: flex;
    flex-direction: column;
    width: 100%;
    max-width: 920px;

    padding-top: 40px;
    margin-bottom: 100px;
  }

  .sub-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .block {
    height: 300px;
    min-width: 400px;
    background: red;
  }
</style>
