<template>
  
    <svg :viewBox="viewBox" :id="id"></svg>
  
</template>

<script>
import * as d3 from "d3";
import _ from "lodash";

export default {
  props: {
    gambars: {
      type: Array,
      required: true,
    },
    id: {
      type: String,
      required: true
    },
    width: {
      default: 720,
      type: Number,
    },
    height: {
      default: 500,
      type: Number,
    },
  },
  data: (vm) => ({
    host: vm.$host,
    padding: 60,
    margin: 60,
    chart: null,
  }),
  computed: {
    rangeX() {
      const width = this.width - this.padding;
      return [0, width];
    },
    rangeY() {
      const height = this.height - this.padding;
      return [0, height];
    },
    viewBox() {
      return `0 0 ${this.width} ${this.height}`;
    },
  },
  watch: {
    gambars(val) {
      if (this.chart != null) this.chart.remove();
      this.render(val);
    },
  },
  methods: {
    render(items) {
      let imageWith = 80;
      let imageHeight = 80;

      const margin = 60;
      const svg_width = this.width;
      const svg_height = this.height;
      const chart_width = svg_width - 2 * margin;
      const chart_height = svg_height - 2 * margin;


      const xScale = d3
        .scaleBand()
        .range([0, chart_width])
        .domain(items.map((s) => s.nama))
        .padding(0.2);


      if (xScale.bandwidth() < imageWith) {
        imageWith = xScale.bandwidth();
        imageHeight = imageWith;
      }


      const bar_height = chart_height - imageHeight;

      const svg = d3
        .select(`svg#${this.id}`)
        .attr("width", svg_width)
        .attr("height", svg_height)
        .attr("viewBox", this.viewBox);

      

      this.chart = svg
        .append("g")
        .attr("transform", `translate(${margin}, ${margin})`);

      const yScale = d3
        .scaleLinear()
        .range([bar_height, 0])
        .domain([0, _.maxBy(items, "bidang").bidang]);

      const color = d3
        .scaleLinear()
        .range(["red", "green"])
        .domain([
          _.maxBy(items, "bidang").bidang,
          0,
        ]);

      this.chart.append("g").call(d3.axisLeft(yScale).ticks(10));

      

      this.chart
        .append("g")
        .attr("transform", `translate(0, ${bar_height})`)
        .call(d3.axisBottom(xScale))
        .selectAll(".tick text")
        .call(wrap, xScale.bandwidth(), imageHeight);

      const barGroups = this.chart.selectAll("rect").data(items).enter();
      barGroups
        .append("rect")
        .attr("fill", (g) => color(g.bidang))
        .attr("x", (g) => xScale(g.nama))
        .attr("y", (g) => yScale(g.bidang))
        .attr("height", (g) => bar_height - yScale(g.bidang))
        .attr("width", xScale.bandwidth())
        .on("mouseenter", function (actual, i) {
          d3.select(this)
            .transition()
            .duration(300)
            .attr("opacity", 0.6)
            .attr("x", (a) => xScale(a.nama) - 5)
            .attr("width", xScale.bandwidth() + 10);
          barGroups
            .append("text")
            .attr("class", "value")
            .attr("x", (a) => xScale(a.nama) + xScale.bandwidth() / 2)
            .attr("y", (a) => yScale(a.bidang) - 20)
            .attr("text-anchor", "middle")
            .text((a, idx) => {
              return idx !== i ? "" : `${a.bidang} bidang`;
            });
        })
        .on("mouseleave", function () {
          d3.selectAll(".bidang").attr("opacity", 1);
          d3.select(this)
            .transition()
            .duration(300)
            .attr("opacity", 1)
            .attr("x", (a) => xScale(a.nama))
            .attr("width", xScale.bandwidth());
          svg.selectAll(".value").remove();
        });

      
      
      barGroups
        .append("image")
        .attr("xlink:href", (g) => `${this.host}upload/${g.foto}`)
        .attr("width", imageWith)
        .attr("height", imageHeight)
        .attr("x", (g) => {
          const center = xScale.bandwidth() / 2 - imageWith / 2;
          return xScale(g.nama) + center;
        })
        .attr("y", bar_height + 9);

      

      barGroups
        .append("text")
        .attr("x", (g) => {
          const center = xScale.bandwidth() / 2;
          return xScale(g.nama) + center;
        })
        .attr("y", (g) => yScale(g.bidang) - 5)
        .attr("text-anchor", "middle")
        .text((a) => {
          return a.bidang;
        });

      svg
        .append("text")
        .attr("class", "label")
        .attr("x", -(bar_height / 2) - margin)
        .attr("y", margin / 2.4)
        .attr("transform", "rotate(-90)")
        .attr("text-anchor", "middle")
        .text("Bidang");
      svg
        .append("text")
        .attr("class", "label")
        .attr("x", chart_width / 2 + margin)
        //.attr("y", chart_height + margin * 1.7)
        .attr("y", svg_height - 10)
        .attr("text-anchor", "middle")
        .text("Petugas Gambar");
      svg
        .append("text")
        .attr("class", "title")
        .attr("x", chart_width / 2 + margin)
        .attr("y", 40)
        .attr("text-anchor", "middle")
        .text("Jumlah bidang yang belum di gambar");

      function wrap(text, width, height) {
        text.each(function () {
          var text = d3.select(this),
            words = text.text().split(/\s+/).reverse(),
            word,
            line = [],
            lineNumber = 0,
            lineHeight = 1.1, // ems
            y = parseFloat(text.attr("y"))  + height + 2,
            dy = parseFloat(text.attr("dy")),
            tspan = text
              .text(null)
              .append("tspan")
              .attr("x", 0)
              .attr("y", y)
              .attr("dy", dy + "em");
            
          while ((word = words.pop())) {
            line.push(word);
            tspan.text(line.join(" "));
            if (tspan.node().getComputedTextLength() > width) {
              line.pop();
              tspan.text(line.join(" "));
              line = [word];
              tspan = text
                .append("tspan")
                .attr("x", 0)
                .attr("y", y)
                .attr("dy", `${++lineNumber * lineHeight + dy}em`)
                .text(word);
            }
          }
        });
      }
    },
  },
};
</script>



