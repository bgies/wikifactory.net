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
            fontSize: '2.2vw'
            lineHeight: '1.6'
            color: '#3c3c3c'
            letterSpacing: ''
            textAlign: 'center'

        h2:
            fontFamily: 'Montserrat'
            fontWeight: 400
            fontSize: '1.8vw'
            color: '#6e6e6e'
            textAlign: 'center'

        h3:
            fontFamily: 'Montserrat'
            fontWeight: 400
            fontSize: '1.4vw'
            color: '#0274B8'
            textAlign: 'center'
            paddingLeft: '1vw'
            paddingRight: '1vw'

        h4: #NAV
            fontFamily: 'Montserrat'
            fontSize: '14px'
            fontWeight: 400
            color: '#3c3c3c'
            marginTop: '5px'
            marginLeft: '25px'

        'h4:hover':
            cursor: 'pointer'
            borderBottom: '1px solid #3c3c3c'

        h5:
            fontFamily: 'Proxima Nova'
            fontWeight: 400
            fontSize: '1.3vw'
            color: '#0274B8'
            textAlign: 'center'
            marginTop: '-0.5vw'

        p:
            fontFamily: 'Proxima Nova'
            fontWeight: 400
            fontSize: '1.5vw'
            color: '#6e6e6e'
            lineHeight: '1.8'
            textAlign: 'center'

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

        # PARTICLES
        canvas:
            display: 'block'
            verticalAlign: 'bottom'

        '#particles-js':
            backgroundColor: '#0274B8' # Wikifactory Blue.
            color: '#ffffff'
            height: '98vh'
            marginBottom: '30px'

        '#particles-js div':
            width: '780px'
            margin: '0 auto'

        '#particles-js img':
            pointerEvents: 'none'
            position: 'absolute'
            top: '30vh'

        '.arrow':
            bottom: '20px'
            width: '50px'
            height: '50px'
            marginLeft: 'calc(50% - 25px)'
            position: 'absolute'

        # WHAT
        '.intro':
            width: '80vw'
            margin: '0 auto'
            paddingTop: '10px'
            paddingBottom: '10px'

        # HOW
        '.nodes':
            width: '100vw'
            height: '30vw'
            marginTop: '30px'
            marginBottom: '40px'
            backgroundImage: 'url("/gfx/nodes.png")'
            position: 'relative'

        'div.nodes-text-one':
            width: '30%'
            position: 'absolute'
            top: '5vw'
            left: '10vw'

        'div.nodes-text-two':
            width: '30%'
            position: 'absolute'
            top: '5vw'
            right: '10vw'

        'div.nodes-text-three':
            width: '30%'
            position: 'absolute'
            top: '18vw'
            left: '20vw'

        'div.nodes-text-four':
            width: '30%'
            position: 'absolute'
            top: '18vw'
            right: '20vw'

        # PROJECTS
        '.projects':
            color: '#3c3c3c'
            marginTop: '80px'
            marginBottom: '10px'

        '.projectimage':
            width: '100vw'
            height: 'auto'
            marginBottom: '30px'

        '.projectimg':
            width: '100%'
            height: 'auto'

        '.device-wrapper':
            width: '65vw'
            margin: '0 auto'
            height: 'auto'

        '.devices':
            width: '100%'
            height: 'auto'

        '.text':
            width: '70vw'
            margin: '0 auto'

        '.small-btn':
            border: '0.2vw solid #0274B8'
            borderRadius: '4px'
            width: '16vw'
            height: '3.6vw'
            margin: '30px auto'
            lineHeight: '0.4'
            transition: 'background-color 0.1s ease-in-out'
            WebkitTransition: 'background-color 0.1s ease-in-out'

        '.small-btn:hover':
            backgroundColor: '#0274B8'
            color: '#fff'
            cursor: 'pointer'

        '.small-btn:hover h3':
            color: '#fff'

        '.medium-btn':
            border: '0.2vw solid #0274B8'
            borderRadius: '4px'
            width: '20vw'
            height: '3.6vw'
            margin: '30px auto'
            lineHeight: '0.4'
            transition: 'background-color 0.1s ease-in-out'
            WebkitTransition: 'background-color 0.1s ease-in-out'

        '.medium-btn:hover':
            backgroundColor: '#0274B8'
            color: '#fff'
            cursor: 'pointer'

        '.medium-btn:hover h3':
            color: '#fff'

        '.large-btn':
            border: '0.2vw solid #0274B8'
            borderRadius: '4px'
            width: '26vw'
            height: '3.6vw'
            margin: '30px auto'
            lineHeight: '0.4'
            transition: 'background-color 0.1s ease-in-out'
            WebkitTransition: 'background-color 0.1s ease-in-out'

        '.large-btn:hover':
            backgroundColor: '#0274B8'
            color: '#fff'
            cursor: 'pointer'

        '.large-btn:hover h3':
            color: '#fff'

        '.partnersimage':
            width: '20vw'
            heihgt: 'auto'
            margin: '0 auto'

        '.partnerimg':
            width: '100%'
            height: 'auto'
            marginTop: '-40px'
            marginBottom: '-30px'

        # TEAM:
        '.team-wrapper':
            width: '90vw'
            height: 'auto'
            margin: '0 auto'
            marginBottom: '40px'

        '.person-card':
            marginTop: '20px'
            width: '20%'
            marginBottom: '30px'
            float: 'left'
            marginLeft: '2.5%'
            marginRight: '2.5%'

        '.person-card h2':
            fontSize: '1.4vw'

        '.person-card h5':
            fontSize: '1.1vw'

        '.person-img':
            paddingLeft: 'auto !important'
            paddingRight: 'auto !important'

        '.avatar':
            borderRadius: '50%'
            width: '12vw'
            height: '12vw'
            marginLeft: '3vw'
            position: 'relative'
            float: 'center'

        '.person-smedia':
            WebkitBoxSizing: 'border-box'
            MozBoxSizing: 'border-box'
            boxSizing: 'border-box'
            clear:'both'
            width: '12vw'
            marginLeft: 'auto'
            marginRight: 'auto'
            height: '30px'
            marginBottom:'10px'
            display:'inline-block'
            float: 'right'

        '.sm-row':
            display: 'inline-block'
            float: 'center'
            marginLeft: '5px'

        '.icon-person':
            WebkitBoxSizing: 'border-box'
            MozBoxSizing: 'border-box'
            boxSizing: 'border-box'
            width: '23px'
            height: 'auto'
            paddingLeft: '4px'
            opacity: 0.5
            float: 'center'

        '.icon-person:hover':
            opacity: 1

        # FOOTER
        '.footer':
            width: '100vw'
            height: '44vw'
            backgroundColor: '#01446c'
            position: 'relative'
            marginTop: '30px'
            clear: 'both'

        '.network-nodes':
            width: '100vw'
            height: '27vw'
            position: 'absolute'
            top: 0
            left: 0
            backgroundImage: 'url("/gfx/nodes-footer.png")'

        '.network-nodes h2':
            color: '#fff'
            marginTop: '40px'
            marginBottom: '20px'

        '.city':
            width: '33.33333%'
            float: 'left'

        '.city h3':
            color: '#fff'
            textAlign: 'center'
            fontFamily: 'Proxima Nova'
            fontSize: '3.0vw'

        '.city h6':
            color: '#fff'
            textAlign: 'center'
            fontFamily: 'Proxima Nova'
            fontWeight: '400'
            fontSize: '1.2vw'
            marginTop: '-2.2vw'

        '.contact':
            width: '8vw'
            height: '2.2vw'
            margin: '0 auto'
            border: '1px solid #fff'
            borderRadius: '4px'
            marginTop: '-1vw'
            transition: 'background-color 0.1s ease-in-out'
            WebkitTransition: 'background-color 0.1s ease-in-out'

        '.contact h5':
            color: '#fff'
            lineHeight: '2.8'
            fontFamily: 'Montserrat'
            fontSize: '1.1vw'

        '.contact:hover':
            backgroundColor: '#fff'
            color: '#fff'
            cursor: 'pointer'

        '.contact:hover h5':
            color: '#0274B8'

        '.new':
            border: '0.2vw solid #fff'
            borderRadius: '4px'
            width: '14vw'
            height: '3.6vw'
            position: 'absolute'
            bottom: '2vw'
            left: 'calc(50% - 7vw)'
            lineHeight: '0.4'
            transition: 'background-color 0.1s ease-in-out'
            WebkitTransition: 'background-color 0.1s ease-in-out'

        '.new h3':
            color: '#fff'
            textAlign: 'center'
            fontFamily: 'Montserrat'
            fontSize: '1.4vw'

        '.new:hover':
            backgroundColor: '#fff'
            color: '#fff'
            cursor: 'pointer'

        '.new:hover h3':
            color: '#0274B8'

        '.nodes-spacer':
            marginTop: '40px'

        '.partners':
            marginTop: '60px'
            marginBottom: '10px'

        '.partners h5':
            fontSize: '1.6vw'
            fontWeight: '600'
            color: '#3c3c3c'
            marginBottom: '10px'

        '.join':
            clear: 'both'
            width: '100vw'
            height: '17vw'
            position: 'absolute'
            bottom: 0
            left: 0

        '.join h2':
           color: '#fff'
            marginTop: '40px'
            marginBottom: '20px'

        '.join p':
            color: 'white'