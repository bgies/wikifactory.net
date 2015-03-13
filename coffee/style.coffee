# Public Domain (-) 2015 The Wikifactory.net Website Authors.
# See the Wikifactory.net UNLICENSE file for details.

module.exports = (api) ->

    api.add
        # Override default box model.
        '*':
            WebkitBoxSizing: 'border-box'
            MozBoxSizing: 'border-box'
            boxSizing: 'border-box'

        # Prevent iOS text size adjust after orientation change, without disabling user zoom.
        html:
            MsTextSizeAdjust: '100%'
            WebkitTextSizeAdjust: '100%'

        # Improve readability when focused and mouse hovered on links in all browsers.
        a:hover:
            outline: 0

        # Global styling.
        body:
            margin: 0
            padding: 0
            # paddingTop: '60px'
            textRendering: 'optimizeLegibility'

        h1:
            fontFamily: 'Montserrat'
            fontWeight: 400
            fontSize: '2.4vw'
            lineHeight: '1.6'
            color: '#3c3c3c'
            letterSpacing: ''
            textAlign: 'center'

        h4: #NAV
            fontFamily: 'Montserrat'
            fontSize: '14px'
            fontWeight: 400
            color: '#6e6e6e'
            marginTop: '5px'
            marginLeft: '25px'

        'h4:hover':
            cursor: 'pointer'
            borderBottom: '1px solid #6e6e6e'

        # 'a::selection':
        #    color: '#fff'

        # '::selection':
        #     background: 'rgba(204,0,0,0.7)'

        # HEADER
         header:
            width: '100%'
            height: '60px'
            backgroundColor: '#fff'
            top: 0
            zIndex: 100
            transition: 'top 0.2s ease-in-out'
            position: 'fixed'
            boxShadow: '0 1px 10px rgba(0, 0, 0, 0.1)'

        '.nav-up':
            top: '-61px'

        'header img':
            width: '42px'
            height: 'auto'
            marginTop: '8px'
            marginLeft: '20px'
            opacity: 0.8

        'header img:hover':
            cursor: 'pointer'
            opacity: 0.9

        # PARTICLES AND TAGLINE
        canvas:
            display: 'block'
            verticalAlign: 'bottom'

        '#particles-js':
            backgroundColor: '#0274B8' # Wikifactory Blue.
            color: '#ffffff'
            height: '98vh'

        '#particles-js div':
            width: '780px'
            margin: '0 auto'

        '#particles-js img':
            pointerEvents: 'none'
            position: 'absolute'
            top: '30vh'

        '.bounce':
            animation: 'bounce 1s infinite'
            WebkitAnimation: 'bounce 1s infinite'

        '.arrow':
            bottom: '20px'
            width: '50px'
            height: '50px'
            marginLeft: 'calc(50% - 25px)'
            position: 'absolute'

        # MAIN WRAPPER
        section:
            paddingTop: '20px'
            width: '85vw'
            margin: '0 auto'

        'nav ul':
            listStyle: 'none'
            float: 'right'
            position: 'absolute'
            right: '25px'
            top: 0

        'nav ul li':
            float: 'right'

        '.cta':
            color: '#0274B8'

        '.cta:hover':
            borderBottom: '1px solid #0274B8'
            cursor: 'pointer'

        '.intro h1':
            marginTop: '50px'
