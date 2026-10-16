// A field of gold dots on a slowly turning globe, behind the hero.
// Atmosphere only: the scene holds no data and takes no input.
import { mount, THREE } from '../threejs.js';

mount('globe', function ({ scene, colour, still }) {
	const root = new THREE.Group();
	scene.add(root);

	root.add(
		new THREE.Points(
			dots(),
			new THREE.PointsMaterial({
				size: 0.032,
				sizeAttenuation: true,
				transparent: true,
				opacity: 0.95,
				color: colour('--three-accent', '#c8962e'),
			}),
		),
	);

	root.add(
		new THREE.LineSegments(
			new THREE.WireframeGeometry(new THREE.SphereGeometry(1, 24, 16)),
			new THREE.LineBasicMaterial({
				transparent: true,
				opacity: 0.16,
				color: colour('--three-line', '#ffffff'),
			}),
		),
	);

	// A slight tilt reads as a globe rather than a ball of noise.
	root.rotation.z = 0.35;
	root.scale.setScalar(1.7);

	// A Fibonacci sphere spreads the dots evenly, with no visible grid.
	function dots() {
		const count = 1200;
		const golden = Math.PI * (3 - Math.sqrt(5));
		const positions = new Float32Array(count * 3);
		for (let i = 0; i < count; i++) {
			const y = 1 - (i / (count - 1)) * 2;
			const r = Math.sqrt(1 - y * y);
			const angle = golden * i;
			positions[i * 3] = Math.cos(angle) * r;
			positions[i * 3 + 1] = y;
			positions[i * 3 + 2] = Math.sin(angle) * r;
		}
		const geometry = new THREE.BufferGeometry();
		geometry.setAttribute('position', new THREE.BufferAttribute(positions, 3));
		return geometry;
	}

	return function (time) {
		if (!still) root.rotation.y = time * 0.06;
	};
});
