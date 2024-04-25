import React, { useRef, useEffect } from 'react';
import * as d3 from 'd3';
import './ResList.css';

function Greph({ nodes, links }) {
    const svgRef = useRef();

    useEffect(() => {
        // Remove existing SVG content
        d3.select(svgRef.current).selectAll('*').remove();

        // Create new SVG
        const svg = d3.select(svgRef.current);

        // Set initial positions for nodes and links
        nodes.forEach((node, index) => {
            node.x = 100 + index * 60; // Adjust X coordinate to set horizontal spacing
            node.y = 100 + index * 50; // Adjust Y coordinate to set vertical spacing
        });

        links.forEach(link => {
            // Calculate midpoints between source and target nodes
            link.x1 = link.source.x;
            link.y1 = link.source.y;
            link.x2 = link.target.x;
            link.y2 = link.target.y;
        });

        const simulation = d3.forceSimulation(nodes)
            .force('link', d3.forceLink(links).id(d => d.id))
            .force('charge', d3.forceManyBody().strength(50))
            .force('center', d3.forceCenter(350, 250));

        const linkForce = d3.forceLink(links)
            .id(d => d.id)
            .distance(700 / nodes.length); // Adjust this value to change link length

        simulation.force('link', linkForce);

        const link = svg.selectAll('line')
            .data(links)
            .enter().append('line')
            .style('stroke', 'blue') // Set color for links
            .attr('x1', d => d.x1)
            .attr('y1', d => d.y1)
            .attr('x2', d => d.x2)
            .attr('y2', d => d.y2);

        const node = svg.selectAll('g')
            .data(nodes)
            .enter().append('g')
            .attr('transform', d => `translate(${d.x},${d.y})`);

        node.append('circle')
            .attr('r', 35)
            .style('fill', 'red') // Set color for nodes
            .style('stroke', 'white');

        node.append('text')
            .attr('dx', 0) // Adjust X offset to position text beside the node
            .attr('dy', -50) // Adjust Y offset for vertical alignment
            .text(d => d.name) // Display node id as text
            .style('fill','white');

        node.call(d3.drag()
            .on('start', dragstarted)
            .on('drag', dragged)
            .on('end', dragended));

        simulation.on('tick', () => {
            link.attr('x1', d => d.source.x)
                .attr('y1', d => d.source.y)
                .attr('x2', d => d.target.x)
                .attr('y2', d => d.target.y);

            node.attr('transform', d => `translate(${d.x},${d.y})`);
        });

        function dragstarted(event, d) {
            if (!event.active) simulation.alphaTarget(0.5).restart();
            d.fx = d.x;
            d.fy = d.y;
        }

        function dragged(event, d) {
            d.fx = event.x;
            d.fy = event.y;
        }

        function dragended(event, d) {
            if (!event.active) simulation.alphaTarget(0);
            d.fx = null;
            d.fy = null;
        }

        return () => simulation.stop();
    }, [nodes, links]);

    return (
        <div>
            <p>Graph Representation</p>
            <svg ref={svgRef} width="900" height="600"></svg>
        </div>
    )
}

export default Greph;
