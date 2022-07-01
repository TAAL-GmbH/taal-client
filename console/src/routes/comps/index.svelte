<script>
  import Button from '../../lib/components/button/index.svelte'
  import ButtonSelect from '../../lib/components/button-select/index.svelte'
  import Dropdown from '../../lib/components/dropdown/index.svelte'
  import TextInput from '../../lib/components/text-input/index.svelte'
  import Checkbox from '../../lib/components/checkbox/index.svelte'
  import Radio from '../../lib/components/radio/index.svelte'
  import Pager from '../../lib/components/pager/index.svelte'
  import Progress from '../../lib/components/progress/index.svelte'
  import Switch from '../../lib/components/switch/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageWithMenu from '../../lib/components/page/template/menu/index.svelte'
  import Row from '../../lib/components/layout/row/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'
  import Modal from '../../lib/components/modal/index.svelte'
  import Popup from '../../lib/components/popup/index.svelte'
  import Text from '../../lib/components/text/index.svelte'
  import { link } from '../../lib/utils/format'

  let popupInputValue

  function onChange(e) {
    console.log('onChange: detail = ', e.detail)
    if (e.detail.name === 'popup-input') {
      popupInputValue = e.detail.value
    }
  }

  let showPopup = false

  function onHidePopup() {
    console.log('onHidePopup')
    showPopup = false
  }
  function onShowPopup() {
    console.log('onShowPopup')
    showPopup = true
  }
</script>

<PageWithMenu>
  <div class="island">
    <Heading value="Components demo" />
    <Spacer h={24} />
    <div class="sub-row">
      <Heading value="Add something" size={2} />
      <Button icon="plus" on:click={onShowPopup}>Add new</Button>
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
    <Spacer h={24} />
    <ButtonSelect
      name="btn-select"
      items={[
        { label: 'One', value: 10 },
        { label: 'Two', value: 20 },
        { label: 'Three', value: 30 },
        { label: 'Filters', value: 77, icon: 'filters' },
        { label: 'List', value: 88, icon: 'view-list' },
        { label: 'Grid', value: 99, iconAfter: 'view-grid' },
      ]}
      on:change={onChange}
    />
    <Spacer h={24} />
    <Pager
      name="bo"
      value={2}
      totalItems={480}
      pageSize={10}
      on:change={onChange}
    />
  </div>
</PageWithMenu>

{#if showPopup}
  <Modal>
    <Popup maxW={480} title="Something" on:close={onHidePopup}>
      <svelte:fragment slot="body">
        <Spacer h={20} />
        <TextInput
          name="popup-input"
          label="Enter something"
          value={popupInputValue}
          required
          on:change={onChange}
        />
        <Spacer h={8} />
        <Text
          size={5}
          html
          value="Examples are often pointed to {link(
            'https://www.example.com'
          )}"
        />
        <Spacer h={80} />
      </svelte:fragment>
      <svelte:fragment slot="footer">
        <Button variant="ghost" size="small" on:click={onHidePopup}
          >Cancel</Button
        >
        <Button size="small" on:click={(e) => console.log('onSubmit')}
          >Submit</Button
        >
      </svelte:fragment>
    </Popup>
  </Modal>
{/if}

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
</style>
