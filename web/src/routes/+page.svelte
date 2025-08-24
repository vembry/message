<!-- <div class="grid h-screen bg-green-300">
	<div class="grid-flow-row gap-4 p-4 bg-gray-400">
		<div class="bg-gray-500 p-4">content</div>
		<div class="bg-gray-500 p-4">content</div>
		<div class="sticky bottom-0 right-0 left-0 gap-4 p-4 bg-green-400">
			<div class="bg-gray-500 p-4">inputs2</div>
		</div>
	</div>
	<div class="absolute bottom-0 right-0 left-0 gap-4 p-4 bg-green-400">
		<div class="bg-gray-500 p-4">inputs</div>
	</div>
</div> -->

<div class="grid h-full bg-stone-200">
	<div class="grid p-2 gap-2">
		{#each { length: 50 } as _, i}
			<div class="rounded-sm bg-stone-400 p-1">message {i}</div>
		{/each}
	</div>
	<div class="sticky bottom-0 left-0 right-0 p-2">
		<div class="flex gap-4">
			<div class="grow">
				<label class="input w-full">
					<input
						type="text"
						placeholder="neutral"
						class="input input-ghost w-full focus:outline-none"
					/>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						fill="none"
						viewBox="0 0 24 24"
						stroke-width="1.5"
						stroke="currentColor"
						class="size-6"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							d="M6 12 3.269 3.125A59.769 59.769 0 0 1 21.485 12 59.768 59.768 0 0 1 3.27 20.875L5.999 12Zm0 0h7.5"
						/>
					</svg>
				</label>
			</div>
			<!-- <div class="rounded-sm bg-stone-400">01</div>
			<div class="rounded-sm bg-stone-400">02</div>
			<div class="rounded-sm bg-stone-400">03</div>
			<div class="col-span-2 rounded-sm bg-stone-400">04</div>
			<div class="rounded-sm bg-stone-400">05</div>
			<div class="rounded-sm bg-stone-400">06</div>
			<div class="col-span-2 rounded-sm bg-stone-400">07</div> -->
		</div>
		<!-- 
		<div class="col-start-auto">
			
		</div> -->
	</div>
</div>
