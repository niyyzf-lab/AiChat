<script setup>
import Sidebar from "./Shell/Sidebar.vue";
import 'vuetify/styles'
import ViewLoader from "./Shell/ViewLoader.vue";

</script>

<template>
  <div id="Body">
    <div id="TitleBar" style="--wails-draggable:drag"></div>
    <div class="Box">
      <Sidebar/>
      <ViewLoader/>
    </div>
  </div>
</template>

<style lang="scss" scoped>
#TitleBar {
  height: 40px;
  width: 100vw;
  cursor: auto;
  position: absolute;
  left: 0;
  top: 5px;
  z-index: 100;
}

#Body {
  position: relative;
  border-radius: 20px;
  //填充视窗
  height: 100vh;
  display: flex;
  flex-direction: column;
  .Box{
    height: 100%;
    display: flex;
  }
}

</style>
