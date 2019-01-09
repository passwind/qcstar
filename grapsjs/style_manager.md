# How does style management work?

style_manager takes responsible of editing the properties of components.

`style_manager`

`index.js` `init` create Sectors & SectorsView

Sectors trigers Sector `buildProperties`, then through `PropertyFactory` build `Property`

SectorsView trigers PropertyView through SectorView. In funtion `initialize`, relative events are set up

```js
this.listenTo(model, 'change:value', this.modelValueChanged);
this.listenTo(model, 'targetUpdated', this.targetUpdated); // how does it work?
```

```js
  modelValueChanged(e, val, opt = {}) {
    const em = this.config.em;
    const model = this.model;
    const value = model.getFullValue();
    const target = this.getTarget();
    const onChange = this.onChange;

    ......

    // here trigger component update instantly
    if (em) {
      em.trigger('component:update', target);
      em.trigger('component:styleUpdate', target);
      em.trigger('component:styleUpdate:' + model.get('property'), target);
    }
  },
```

**Question:** When `target` component style is changed, then `targetUpdated()` is triggered?

